package public

import (
	"context"
	"moke/internal/auth/service/utils"

	"github.com/pkg/errors"

	pb "moke/proto/gen/auth/api"
)

func (s *Service) Authenticate(_ context.Context, request *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	if data, err := s.db.LoadOrCreateUid(request.Username); err != nil {
		return nil, errors.Wrap(ErrGetIDFailure, err.Error())
	} else if access, err := utils.CreatJwt(data.GetUid(), utils.TokenTypeAccess, s.jwtSecret); err != nil {
		return nil, errors.Wrap(ErrGenerateJwtFailure, err.Error())
	} else if refresh, err := utils.CreatJwt(data.GetUid(), utils.TokenTypeRefresh, s.jwtSecret); err != nil {
		return nil, errors.Wrap(ErrGenerateJwtFailure, err.Error())
	} else {
		return &pb.AuthenticateResponse{
			AccessToken:  access,
			RefreshToken: refresh,
		}, nil
	}
}

func (s *Service) ValidateToken(_ context.Context, request *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	if uid, err := utils.ParseToken(request.AccessToken, utils.TokenTypeAccess, s.jwtSecret); err != nil {
		return nil, errors.Wrap(ErrParseJwtTokenFailure, err.Error())
	} else {
		return &pb.ValidateTokenResponse{
			Uid: uid,
		}, nil
	}
}

func (s *Service) RefreshToken(_ context.Context, request *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	if uid, err := utils.ParseToken(request.RefreshToken, utils.TokenTypeRefresh, s.jwtSecret); err != nil {
		return nil, errors.Wrap(ErrParseJwtTokenFailure, err.Error())
	} else if access, err := utils.CreatJwt(uid, utils.TokenTypeAccess, s.jwtSecret); err != nil {
		return nil, errors.Wrap(ErrGenerateJwtFailure, err.Error())
	} else if refresh, err := utils.CreatJwt(uid, utils.TokenTypeRefresh, s.jwtSecret); err != nil {
		return nil, errors.Wrap(ErrGenerateJwtFailure, err.Error())
	} else {
		return &pb.RefreshTokenResponse{
			AccessToken:  access,
			RefreshToken: refresh,
		}, nil
	}
}
