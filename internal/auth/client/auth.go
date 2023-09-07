package client

import (
	"context"
	"fmt"

	"github.com/abiosoft/ishell"
	mm "github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/gstones/moke-kit/utility/ugrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/gstones/moke-kit/utility/cshell"

	pb2 "moke/proto/gen/auth/api"
)

type AuthClient struct {
	client pb2.AuthServiceClient
	cmd    *ishell.Cmd
	conn   *grpc.ClientConn
}

func NewAuthClient(host string) (*AuthClient, error) {
	if conn, err := ugrpc.DialWithOptions(host, false); err != nil {
		return nil, err
	} else {
		cmd := &ishell.Cmd{
			Name:    "auth",
			Help:    "auth interactive",
			Aliases: []string{"A"},
		}
		p := &AuthClient{
			client: pb2.NewAuthServiceClient(conn),
			cmd:    cmd,
			conn:   conn,
		}
		p.initSubShells()
		return p, nil
	}
}

func (p *AuthClient) GetCmd() *ishell.Cmd {
	return p.cmd
}

func (p *AuthClient) Close() error {
	return p.conn.Close()
}

func (p *AuthClient) initSubShells() {
	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "auth",
		Help:    "authorize interactive",
		Aliases: []string{"A"},
		Func:    p.auth,
	})
	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "validate",
		Help:    "validate interactive",
		Aliases: []string{"V"},
		Func:    p.validate,
	})
	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "refresh",
		Help:    "refresh interactive",
		Aliases: []string{"R"},
		Func:    p.refresh,
	})

}

func (p *AuthClient) auth(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	cshell.Info(c, "Enter username...")
	msg := cshell.ReadLine(c, "username: ")
	req := &pb2.AuthenticateRequest{
		Username: msg,
		AppId:    "test",
		Platform: pb2.AuthenticateRequest_WECHAT,
	}

	md := metadata.Pairs("authorization", fmt.Sprintf("%s %v", "bearer", "test"))
	ctx := mm.MD(md).ToOutgoing(context.Background())
	if response, err := p.client.Authenticate(ctx, req); err != nil {
		cshell.Warn(c, err)
	} else {
		cshell.Infof(c, "Response: access %s", response.AccessToken)
		cshell.Infof(c, "Response: refresh %s", response.RefreshToken)
	}
}

func (p *AuthClient) validate(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	cshell.Info(c, "Enter token...")
	msg := cshell.ReadLine(c, "token: ")
	req := &pb2.ValidateTokenRequest{
		AccessToken: msg,
	}
	//md := metadata.Pairs("authorization", fmt.Sprintf("%s %v", "bearer", "test"))
	//ctx := mm.MD(md).ToOutgoing(context.Background())
	if response, err := p.client.ValidateToken(context.TODO(), req); err != nil {
		cshell.Warn(c, err)
	} else {
		cshell.Infof(c, "Response: %s", response)
	}
}

func (p *AuthClient) refresh(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	cshell.Info(c, "Enter refresh token...")
	msg := cshell.ReadLine(c, "refresh token: ")

	req := &pb2.RefreshTokenRequest{
		RefreshToken: msg,
	}
	//md := metadata.Pairs("authorization", fmt.Sprintf("%s %v", "bearer", "test"))
	//ctx := mm.MD(md).ToOutgoing(context.Background())
	if response, err := p.client.RefreshToken(context.TODO(), req); err != nil {
		cshell.Warn(c, err)
	} else {
		cshell.Infof(c, "Response: refresh %s", response.RefreshToken)
		cshell.Infof(c, "Response: access %s", response.AccessToken)
	}
}
