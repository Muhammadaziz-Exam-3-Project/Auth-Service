package repositoty

import (
	"Github.com/LocalEats/Authentication-Service/internal/configs/logger"
	"context"

	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

	q "github.com/Masterminds/squirrel"

	pb "Github.com/LocalEats/Authentication-Service/gen-proto/auth"
)

type AuthRepository struct {
	DB      *sql.DB
	Builder q.StatementBuilderType
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		DB:      db,
		Builder: q.StatementBuilder.PlaceholderFormat(q.Dollar),
	}
}

func (s *AuthRepository) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO users (id, username, email, password_hash,full_name, user_type)values($1, $2, $3, $4)`

	password, err := HashPassword(req.Password)
	if err != nil {
		log.Error("Failed to hash password")
		return nil, err
	}
	id := uuid.NewString()

	req.Password = string(password)

	_, err = s.DB.ExecContext(ctx, query, id, req.Username, req.Email, password, req.FullName, req.UserType)
	if err != nil {
		log.Error("Failed to insert user")
		return nil, err
	}
	resp := &pb.User{
		Id:        id,
		Username:  req.Username,
		Email:     req.Email,
		FullName:  req.FullName,
		UserType:  req.UserType,
		CreatedAt: cast.ToString(time.Now()),
	}

	log.Info("Successfully registered user")
	return &pb.RegisterResponse{User: resp}, nil
}

func (s *AuthRepository) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	res := pb.LoginResponse{}
	query :=
		`
			SELECT password_hash from users WHERE email = $1
		`

	var password_hash string
	err := s.DB.QueryRowContext(ctx, query, req.Email).Scan(&password_hash)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("Failed to find user")
			return nil, err
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(req.Password))
	if err != nil {
		log.Error("Failed to find user")
		return &res, fmt.Errorf("invalid password")
	}

	log.Info("Successfully logged in")
	return &res, nil
}

func (s *AuthRepository) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {

	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	var resp pb.GetProfileResponse

	query := "select * from users where id = $1"

	resp.User.CreatedAt = cast.ToString(time.Now())
	resp.User.UpdatedAt = cast.ToString(time.Now())

	row := s.DB.QueryRowContext(ctx, query, req.ID)
	err = row.Scan(&resp.User.Id, &resp.User.Username, &resp.User.Email, &resp.User.FullName, &resp.User.UserType, &resp.User.Address, &resp.User.PhoneNumber, &resp.User.CreatedAt, &resp.User.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("User not found")
			return nil, status.Error(codes.NotFound, "User not found")
		}
		return nil, err
	}

	log.Info("Successful get profile")
	return &resp, nil
}

func (s *AuthRepository) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {

	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	updatedAt := cast.ToString(time.Now())

	query, args, err := s.Builder.Update("users").
		Set("full_name", req.FullName).
		Set("address", req.Address).
		Set("phone_number", req.PhoneNumber).
		Set("updated_at", updatedAt).
		Where(q.Eq{"user_id": req.ID}).
		Suffix("RETURNING username, email, user_type").
		ToSql()

	if err != nil {
		log.Error("Failed to build query")
		return nil, err
	}
	row := s.DB.QueryRowContext(ctx, query, args...)
	var username, email, userType string
	err = row.Scan(&username, &email, &userType)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("User not found")
			return nil, status.Error(codes.NotFound, "User not found")
		}
		return nil, err
	}
	resp := &pb.User{
		Id:        req.ID,
		Username:  username,
		Email:     email,
		FullName:  req.FullName,
		UserType:  userType,
		CreatedAt: cast.ToString(time.Now()),
		UpdatedAt: cast.ToString(time.Now()),
	}

	log.Info("Successfully updated user")
	return &pb.UpdateProfileResponse{User: resp}, nil
}

func (s *AuthRepository) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	query, args, err := s.Builder.Delete("users").
		Where(q.Eq{"email": req.Email}).
		ToSql()
	if err != nil {
		log.Error("Failed to build query")
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	result, err := s.DB.ExecContext(ctx, query, args...)
	if err != nil {
		log.Error("Failed to reset password")
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error("Failed to get rows affected")
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	if rowsAffected == 0 {
		log.Error("No rows affected")
		return nil, status.Error(codes.NotFound, "User not found")
	}

	return &pb.ResetPasswordResponse{Message: "Password successfully updated"}, nil
}

func (s *AuthRepository) RefreshToken(ctx context.Context, id string, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	query := "update token set refresh_token = $1 where id = $2"
	_, err = s.DB.ExecContext(ctx, query, req.RefreshToken, id)
	if err != nil {
		log.Error("Failed to refresh token")
		return nil, err
	}
	return &pb.RefreshTokenResponse{RefreshToken: req.RefreshToken}, nil
}

func (s *AuthRepository) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return &pb.LogoutResponse{Message: "User successfully deleted"}, nil
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
