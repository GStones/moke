package service

import (
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/qfx"
	"github.com/gstones/moke-kit/orm/nosql/diface"
	"github.com/gstones/moke-kit/orm/pkg/nfx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"moke/internal/buddy/db"
	"moke/internal/buddy/pkg/bfx"
	pb "moke/proto/gen/buddy/api"
)

type Service struct {
	logger     *zap.Logger
	db         *db.Database
	mq         miface.MessageQueue
	maxInviter int32
	maxBuddies int32
	maxBlocked int32
}

func NewService(
	l *zap.Logger,
	coll diface.ICollection,
	mq miface.MessageQueue,
	setting bfx.BuddySettings,
) (result *Service, err error) {
	result = &Service{
		logger:     l,
		db:         db.OpenDatabase(l, coll),
		mq:         mq,
		maxBuddies: setting.BuddyMaxCount,
		maxBlocked: setting.BlockedMaxCount,
		maxInviter: setting.InviterMaxCount,
	}
	return
}

func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	pb.RegisterBuddyServiceServer(server.GrpcServer(), s)
	return nil
}

var Module = fx.Provide(
	func(
		l *zap.Logger,
		dProvider nfx.DocumentStoreParams,
		mqParams qfx.MessageQueueParams,
		setting bfx.BuddySettings,
	) (out sfx.GrpcServiceResult, err error) {
		if coll, e := dProvider.DriverProvider.OpenDbDriver(setting.Name); e != nil {
			err = e
		} else if s, e := NewService(l, coll, mqParams.MessageQueue, setting); e != nil {
			err = e
		} else {
			out.GrpcService = s
		}
		return
	},
)
