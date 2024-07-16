package service

import (
	pb "Github.com/LocalEats/Authentication-Service/gen-proto/auth"
	"Github.com/LocalEats/Authentication-Service/internal/configs/logger"
	"context"
)

func (a *AuthService) CreateKitchen(ctx context.Context, req *pb.CreateKitchenRequest) (*pb.CreateKitchenResponse, error) {
	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	log.Info("CreateKitchen called")
	return a.Auth.CreateKitchen(ctx, req)
}

func (a *AuthService) UpdateKitchen(ctx context.Context, req *pb.UpdateKitchenRequest) (*pb.UpdateKitchenResponse, error) {
	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	log.Info("UpdateKitchen called")
	return a.Auth.UpdateKitchen(ctx, req)
}

func (a *AuthService) GetKitchen(ctx context.Context, req *pb.GetKitchenRequest) (*pb.GetKitchenResponse, error) {
	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	log.Info("GetKitchen called")
	return a.Auth.GetKitchen(ctx, req)
}

func (a *AuthService) ListKitchens(ctx context.Context, req *pb.ListKitchensRequest) (*pb.ListKitchensResponse, error) {
	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	log.Info("ListKitchens called")
	return a.Auth.ListKitchens(ctx, req)
}

func (a *AuthService) SearchKitchen(ctx context.Context, req *pb.SearchKitchensRequest) (*pb.SearchKitchensResponse, error) {
	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	log.Info("SearchKitchens called")
	return a.Auth.SearchKitchen(ctx, req)
}
