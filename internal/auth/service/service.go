package public

import (
	"github.com/gstones/moke-kit/fxmain/pkg/mfx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"
	"go.uber.org/fx"
	"go.uber.org/zap"
	mfx2 "moke/internal/auth/pkg/mfx"
	pb "moke/proto/gen/auth/api"
)

type Service struct {
	appId      string
	logger     *zap.Logger
	deployment string
	jwtSecret  string
	url        string
}

func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	pb.RegisterAuthServiceServer(server.GrpcServer(), s)
	return nil
}

func NewService(
	l *zap.Logger,
	deployment string,
	jwtSecret string,
	url string,
) (result *Service, err error) {
	result = &Service{
		logger:     l,
		deployment: deployment,
		jwtSecret:  jwtSecret,
		url:        url,
	}
	return
}

var ServiceModule = fx.Provide(
	func(
		l *zap.Logger,
		aSetting mfx.AppParams,
		sSetting mfx2.AuthSettingParams,
	) (out sfx.GrpcServiceResult, err error) {
		if svc, e := NewService(
			l,
			aSetting.Deployment,
			sSetting.JwtTokenSecret,
			sSetting.AuthUrl,
		); e != nil {
			err = e
		} else {
			out.GrpcService = svc
		}
		return
	},
)
