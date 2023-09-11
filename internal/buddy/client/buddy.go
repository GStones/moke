package client

import (
	"context"
	"fmt"

	"github.com/abiosoft/ishell"
	mm "github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/gstones/moke-kit/utility/cshell"
	"github.com/gstones/moke-kit/utility/ugrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	buddy "moke/proto/gen/buddy/api"
)

type BuddyClient struct {
	client buddy.BuddyServiceClient
	cmd    *ishell.Cmd
	conn   *grpc.ClientConn
}

func NewBuddyClient(host string) (*BuddyClient, error) {
	if conn, err := ugrpc.DialWithOptions(host, false); err != nil {
		return nil, err
	} else {
		cmd := &ishell.Cmd{
			Name:    "buddy",
			Help:    "buddy interactive",
			Aliases: []string{"B"},
		}
		p := &BuddyClient{
			client: buddy.NewBuddyServiceClient(conn),
			cmd:    cmd,
			conn:   conn,
		}
		p.initSubShells()
		return p, nil
	}
}

func (p *BuddyClient) GetCmd() *ishell.Cmd {
	return p.cmd
}

func (p *BuddyClient) Close() error {
	return p.conn.Close()
}

func (p *BuddyClient) initSubShells() {
	p.cmd.AddCmd(&ishell.Cmd{
		Name:    "add",
		Help:    "add buddy",
		Aliases: []string{"A"},
		Func:    p.add,
	})

}

func (p *BuddyClient) add(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	cshell.Info(c, "Enter buddy name...")
	msg := cshell.ReadLine(c, "buddy name: ")
	req := &buddy.AddBuddyRequest{
		Uid:     []string{msg},
		ReqInfo: "test",
	}

	md := metadata.Pairs("authorization", fmt.Sprintf("%s %v", "bearer", "test"))
	ctx := mm.MD(md).ToOutgoing(context.Background())

	if response, err := p.client.AddBuddy(ctx, req); err != nil {
		cshell.Warn(c, err)
	} else {
		cshell.Infof(c, "Response: %s", response)
	}

}
