package service

import (
	pb "Github.com/LocalEats/Authentication-Service/gen-proto/auth"
	"Github.com/LocalEats/Authentication-Service/internal/configs/logger"
	r "Github.com/LocalEats/Authentication-Service/internal/repositoty"
	"context"
)

type AuthService struct {
	Auth r.AuthRepository
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(auth r.AuthRepository) *AuthService {
	return &AuthService{
		Auth: auth,
	}
}

func (a *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	log.Info("Login called")
	return a.Auth.Login(ctx, req)
}

func (a *AuthService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	log.Info("Logout called")
	return a.Auth.Logout(ctx, req)
}

func (a *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	log.Info("Register called")
	return a.Auth.Register(ctx, req)
}

func (a *AuthService) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	log.Info("ResetPassword called")
	return a.Auth.ResetPassword(ctx, req)
}
