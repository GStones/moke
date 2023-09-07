package public

import (
	"github.com/gstones/moke-kit/orm/nosql/diface"
	"github.com/gstones/moke-kit/orm/pkg/nfx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"
	"go.uber.org/fx"
	"go.uber.org/zap"

	mfx2 "moke/internal/auth/pkg/mfx"
	"moke/internal/auth/service/db"
	pb "moke/proto/gen/auth/api"
)

type Service struct {
	logger    *zap.Logger
	jwtSecret string
	url       string
	db        *db.Database
}

func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	pb.RegisterAuthServiceServer(server.GrpcServer(), s)
	return nil
}

func NewService(
	l *zap.Logger,
	jwtSecret string,
	url string,
	coll diface.ICollection,
) (result *Service, err error) {
	result = &Service{
		logger:    l,
		jwtSecret: jwtSecret,
		url:       url,
		db:        db.OpenDatabase(l, coll),
	}
	return
}

var ServiceModule = fx.Provide(
	func(
		l *zap.Logger,
		sSetting mfx2.AuthSettingParams,
		dbProvider nfx.DocumentStoreParams,
	) (out sfx.GrpcServiceResult, err error) {
		if coll, e := dbProvider.DriverProvider.OpenDbDriver(sSetting.AuthStoreName); e != nil {
			err = e
		} else {
			if svc, e := NewService(
				l,
				sSetting.JwtTokenSecret,
				sSetting.AuthUrl,
				coll,
			); e != nil {
				err = e
			} else {
				out.GrpcService = svc
			}
		}
		return
	},
)
