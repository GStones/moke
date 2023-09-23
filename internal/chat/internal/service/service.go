package service

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/gstones/moke-kit/mq/miface"
	"github.com/gstones/moke-kit/mq/pkg/qfx"
	"github.com/gstones/moke-kit/orm/pkg/nfx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"
	"go.uber.org/zap"

	"go.uber.org/fx"

	"moke/internal/chat/pkg/cfx"
	pb "moke/proto/gen/chat/api"
)

type Service struct {
	logger *zap.Logger
	mq     miface.MessageQueue

	redis *redis.Client

	chatInterval time.Duration
}

func NewService(
	l *zap.Logger,
	rClient *redis.Client,
	mq miface.MessageQueue,
	chatInterval int,
) (result *Service, err error) {
	result = &Service{
		logger:       l,
		redis:        rClient,
		mq:           mq,
		chatInterval: time.Duration(chatInterval) * time.Second,
	}
	return
}

func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	pb.RegisterChatServiceServer(server.GrpcServer(), s)
	return nil
}

var ChatService = fx.Provide(
	func(
		l *zap.Logger,
		setting cfx.ChatSettingParams,
		mqParams qfx.MessageQueueParams,
		redisParams nfx.RedisParams,

	) (out sfx.GrpcServiceResult, err error) {
		if s, err := NewService(
			l,
			redisParams.Redis,
			mqParams.MessageQueue,
			setting.ChatInterval,
		); err != nil {
			return out, err
		} else {
			out.GrpcService = s
		}
		return
	},
)
