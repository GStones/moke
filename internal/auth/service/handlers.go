package public

import (
	"context"
	pb "moke/proto/gen/auth/api"
)

func (s *Service) Authenticate(ctx context.Context, request *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ValidateToken(ctx context.Context, request *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) RefreshToken(ctx context.Context, request *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}
