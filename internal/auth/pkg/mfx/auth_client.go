package mfx

import (
	"github.com/gstones/moke-kit/utility/ugrpc"
	"go.uber.org/fx"

	pb "moke/proto/gen/auth/api"
)

type AuthClientParams struct {
	fx.In

	AuthClient pb.AuthServiceClient `name:"AuthClient"`
}

type AuthClientResult struct {
	fx.Out

	AuthClient pb.AuthServiceClient `name:"AuthClient"`
}

func NewAuthClient(host string) (pb.AuthServiceClient, error) {
	if conn, err := ugrpc.DialWithOptions(host, false); err != nil {
		return nil, err
	} else {
		return pb.NewAuthServiceClient(conn), nil
	}
}

var AuthClientModule = fx.Provide(
	func(
		setting AuthSettingParams,
	) (out AuthClientResult, err error) {
		if cli, e := NewAuthClient(setting.AuthUrl); e != nil {
			err = e
		} else {
			out.AuthClient = cli
		}
		return
	},
)
