package mfx

import (
	"context"

	"github.com/gstones/moke-kit/server/pkg/sfx"
	"go.uber.org/fx"
	"go.uber.org/zap"

	pb "moke/proto/gen/auth/api"
)

// Author is auth for grpc middleware
type Author struct {
	client pb.AuthServiceClient
}

// Auth will  auth  every grpc request
func (d *Author) Auth(token string) (string, error) {
	if resp, err := d.client.ValidateToken(context.TODO(), &pb.ValidateTokenRequest{
		AccessToken: token,
	}); err != nil {
		return "", err
	} else {
		return resp.GetUid(), err
	}
}

// AuthCheckModule is the module for grpc middleware
var AuthCheckModule = fx.Provide(
	func(
		l *zap.Logger,
		params AuthClientParams,
	) (out sfx.AuthServiceResult, err error) {
		out.AuthService = &Author{
			client: params.AuthClient,
		}
		return
	},
)
