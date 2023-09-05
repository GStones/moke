package client

import (
	"context"
	"fmt"

	"github.com/abiosoft/ishell"
	mm "github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "github.com/gstones/moke-kit/demo/api/gen/demo/api"
	"github.com/gstones/moke-kit/utility/cshell"
)

type AuthClient struct {
	client pb.DemoClient
	cmd    *ishell.Cmd
}

func NewAuthClient(conn *grpc.ClientConn) *AuthClient {
	cmd := &ishell.Cmd{
		Name:    "auth",
		Help:    "auth interactive",
		Aliases: []string{"A"},
	}
	p := &AuthClient{
		client: pb.NewDemoClient(conn),
		cmd:    cmd,
	}
	p.initSubShells()
	return p
}

func (p *AuthClient) GetCmd() *ishell.Cmd {
	return p.cmd
}

func (p *AuthClient) initSubShells() {
	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "auth",
		Help:    "auth interactive",
		Aliases: []string{"A"},
		Func:    p.auth,
	})
	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "watch",
		Help:    "watch topic",
		Aliases: []string{"w"},
		Func:    p.watch,
	})

}

func (p *AuthClient) auth(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	cshell.Info(c, "Enter say hi message...")
	msg := cshell.ReadLine(c, "message: ")

	md := metadata.Pairs("authorization", fmt.Sprintf("%s %v", "bearer", "test"))
	ctx := mm.MD(md).ToOutgoing(context.Background())
	if response, err := p.client.Hi(ctx, &pb.HiRequest{
		Uid:     "10000",
		Message: msg,
	}); err != nil {
		cshell.Warn(c, err)
	} else {
		cshell.Infof(c, "Response: %s", response.Message)
	}
}

func (p *AuthClient) watch(c *ishell.Context) {

	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	cshell.Info(c, "Enter watch topic...")
	topic := cshell.ReadLine(c, "topic: ")

	md := metadata.Pairs("authorization", fmt.Sprintf("%s %v", "bearer", "test"))
	ctx := mm.MD(md).ToOutgoing(context.Background())
	if stream, err := p.client.Watch(ctx, &pb.WatchRequest{
		Topic: topic,
	}); err != nil {
		cshell.Warn(c, err)
	} else {
		for {
			if response, err := stream.Recv(); err != nil {
				cshell.Warn(c, err)
				break
			} else {
				cshell.Infof(c, "Response: %s \r\n", response.Message)
			}
		}
	}
}
