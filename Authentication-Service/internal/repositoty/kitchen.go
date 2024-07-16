package repositoty

import (
	pb "Github.com/LocalEats/Authentication-Service/gen-proto/auth"
	"Github.com/LocalEats/Authentication-Service/internal/configs/logger"
	"context"
	q "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"go.uber.org/zap"

	"time"
)

func (s *AuthRepository) CreateKitchen(ctx context.Context, k *pb.CreateKitchenRequest) (*pb.CreateKitchenResponse, error) {
	k.Kitchen.Id = uuid.NewString()
	k.Kitchen.CreatedAt = cast.ToString(time.Now())

	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	query := `insert into kitchens (id, owner_id,name, description,cuisine_type,address,phone_number, created_at)
    values ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err = s.DB.ExecContext(ctx, query, k.Kitchen.Id, k.Kitchen.OwnerId, k.Kitchen.Name, k.Kitchen.Description, k.Kitchen.CuisineType, k.Kitchen.Address, k.Kitchen.PhoneNumber, k.Kitchen.CreatedAt)
	if err != nil {
		log.Error("Failed to create kitchen", zap.Error(err))
		return nil, err
	}

	resp := &pb.CreateKitchenResponse{
		Kitchen: &pb.Kitchen{
			Id:          k.Kitchen.Id,
			OwnerId:     k.Kitchen.OwnerId,
			Name:        k.Kitchen.Name,
			Description: k.Kitchen.Description,
			CuisineType: k.Kitchen.CuisineType,
			Address:     k.Kitchen.Address,
			PhoneNumber: k.Kitchen.PhoneNumber,
			Rating:      0,
			CreatedAt:   k.Kitchen.CreatedAt,
		},
	}

	log.Info("Create Kitchen", zap.Any("kitchen", resp))
	return resp, nil
}

func (s *AuthRepository) UpdateKitchen(ctx context.Context, kitchen *pb.UpdateKitchenRequest) (*pb.UpdateKitchenResponse, error) {

	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	query, args, err := s.Builder.Update("kitchens").
		Set("name", kitchen.Name).
		Set("description", kitchen.Description).
		Where(q.Eq{"kitchen_id": kitchen.KitchenId}).
		ToSql()
	if err != nil {
		log.Error("Failed to build update kitchen", zap.Error(err))
		return nil, err
	}

	var owner string
	var cuisineType string
	var rating float32

	row := s.DB.QueryRowContext(ctx, query, args...)
	err = row.Scan(&owner, &cuisineType, &rating)
	if err != nil {
		log.Error("Failed to update kitchen", zap.Error(err))
		return nil, err
	}

	resp := &pb.UpdateKitchenResponse{
		KitchenId:   kitchen.KitchenId,
		OwnerId:     owner,
		Name:        kitchen.Name,
		Description: kitchen.Description,
		CuisineType: cuisineType,
		Rating:      rating,
		UpdatedAt:   cast.ToString(time.Now()),
	}
	log.Info("UpdateKitchen", zap.Any("kitchen", resp))
	return resp, nil
}

func (s *AuthRepository) GetKitchen(ctx context.Context, req *pb.GetKitchenRequest) (*pb.GetKitchenResponse, error) {

	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	query, args, err := s.Builder.Select(
		"owner_id",
		"name",
		"description",
		"cuisine_type",
		"address",
		"phone_number",
		"rating",
		"total_orders",
		"created_at",
		"updated_at").
		From("kitchens").
		Where(q.Eq{"kitchen_id": req.KitchenId}).
		ToSql()
	if err != nil {
		log.Error("Failed to build update kitchen", zap.Error(err))
		return nil, err
	}

	var resp pb.GetKitchenResponse

	row := s.DB.QueryRowContext(ctx, query, args...)
	err = row.Scan(
		&resp.Kitchen.Id,
		&resp.Kitchen.OwnerId,
		&resp.Kitchen.Name,
		&resp.Kitchen.Description,
		&resp.Kitchen.CuisineType,
		&resp.Kitchen.Address,
		&resp.Kitchen.PhoneNumber,
		&resp.Kitchen.Rating,
		&resp.Kitchen.TotalOrders,
		&resp.Kitchen.CreatedAt,
		&resp.Kitchen.UpdatedAt)

	if err != nil {
		log.Error("Failed to get kitchen", zap.Error(err))
		return nil, err
	}

	log.Info("GetKitchen", zap.Any("kitchen", resp))
	return &resp, nil
}

func (s *AuthRepository) ListKitchens(ctx context.Context, req *pb.ListKitchensRequest) (*pb.ListKitchensResponse, error) {

	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	var total int32
	countQuery, countArgs, err := s.Builder.Select("COUNT(*)").
		From("kitchens").
		ToSql()
	if err != nil {
		log.Error("Failed to build update kitchens", zap.Error(err))
		return nil, err
	}

	err = s.DB.QueryRowContext(ctx, countQuery, countArgs...).Scan(&total)
	if err != nil {
		log.Error("Failed to list kitchens", zap.Error(err))
		return nil, err
	}

	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}

	totalPages := (total + limit - 1) / limit
	page := req.Page
	if page <= 0 {
		page = 1
	}
	if page > totalPages {
		page = totalPages
	}

	offset := (page - 1) * limit

	// Main query to fetch kitchens
	query, args, err := s.Builder.Select(
		"kitchen_id",
		"name",
		"cuisine_type",
		"rating",
		"total_orders").
		From("kitchens").
		OrderBy("rating DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		log.Error("Failed to build update kitchens", zap.Error(err))
		return nil, err
	}

	rows, err := s.DB.QueryContext(ctx, query, args...)
	if err != nil {
		log.Error("Failed to list kitchens", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var kitchens []*pb.Kitchen

	for rows.Next() {
		kitchen := &pb.Kitchen{}
		err = rows.Scan(
			&kitchen.Id,
			&kitchen.Name,
			&kitchen.CuisineType,
			&kitchen.Rating,
			&kitchen.TotalOrders,
		)
		if err != nil {
			log.Error("Failed to list kitchens", zap.Error(err))
			return nil, err
		}
		kitchens = append(kitchens, kitchen)
	}

	if err = rows.Err(); err != nil {
		log.Error("Failed to list kitchens", zap.Error(err))
		return nil, err
	}

	log.Info("List Kitchens", zap.Any("kitchens", kitchens))

	return &pb.ListKitchensResponse{
		Kitchens: kitchens,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}

func (s *AuthRepository) SearchKitchen(ctx context.Context, req *pb.SearchKitchensRequest) (*pb.SearchKitchensResponse, error) {
	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	var total int32
	countQuery, countArgs, err := s.Builder.Select("COUNT(*)").
		From("kitchens").
		Where(q.Like{"name": "%" + req.Name + "%"}).
		ToSql()
	if err != nil {
		log.Error("Failed to build count query", zap.Error(err))
		return nil, err
	}

	err = s.DB.QueryRowContext(ctx, countQuery, countArgs...).Scan(&total)
	if err != nil {
		log.Error("Failed to count kitchens", zap.Error(err))
		return nil, err
	}

	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}

	totalPages := (total + limit - 1) / limit
	page := req.Page
	if page <= 0 {
		page = 1
	}
	if page > totalPages {
		page = totalPages
	}

	offset := (page - 1) * limit

	query, args, err := s.Builder.Select(
		"kitchen_id",
		"name",
		"cuisine_type",
		"rating",
		"total_orders").
		From("kitchens").
		Where(q.Like{"name": "%" + req.Name + "%"}).
		OrderBy("rating DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		log.Error("Failed to search kitchens", zap.Error(err))
		return nil, err
	}

	rows, err := s.DB.QueryContext(ctx, query, args...)
	if err != nil {
		log.Error("Failed to search kitchens", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var kitchens []*pb.KitchenList

	for rows.Next() {
		kitchen := &pb.KitchenList{}
		err = rows.Scan(
			&kitchen.Id,
			&kitchen.Name,
			&kitchen.CuisineType,
			&kitchen.Rating,
			&kitchen.TotalOrders,
		)
		if err != nil {
			log.Error("Failed to scan rows", zap.Error(err))
			return nil, err
		}
		kitchens = append(kitchens, kitchen)
	}

	if err = rows.Err(); err != nil {
		log.Error("Failed to list kitchens", zap.Error(err))
		return nil, err
	}

	return &pb.SearchKitchensResponse{
		Kitchens: kitchens,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}
