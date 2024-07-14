package repositoty

import (
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "Authentication-Service/gen-proto/auth"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (s *AuthRepository) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func (s *AuthRepository) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func (s *AuthRepository) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}

func (s *AuthRepository) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}

func (s *AuthRepository) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}

func (s *AuthRepository) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}

func (s *AuthRepository) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}

func (s *AuthRepository) CreateKitchen(ctx context.Context, req *pb.CreateKitchenRequest) (*pb.CreateKitchenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKitchen not implemented")
}

func (s *AuthRepository) UpdateKitchen(ctx context.Context, req *pb.UpdateKitchenRequest) (*pb.UpdateKitchenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKitchen not implemented")
}

func (s *AuthRepository) GetKitchen(ctx context.Context, req *pb.GetKitchenRequest) (*pb.GetKitchenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKitchen not implemented")
}

func (s *AuthRepository) ListKitchens(ctx context.Context, req *pb.ListKitchensRequest) (*pb.ListKitchensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKitchens not implemented")
}

func (s *AuthRepository) SearchKitchens(ctx context.Context, req *pb.SearchKitchensRequest) (*pb.SearchKitchensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchKitchens not implemented")
}
