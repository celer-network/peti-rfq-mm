package rfq

import (
	"context"

	"github.com/celer-network/rfq-mm/sdk/service/rfq/proto"
	"google.golang.org/grpc"
)

type Client struct {
	proto.MMApiClient
	server string
	conn   *grpc.ClientConn
}

type UserClient struct {
	proto.UserApiClient
	server string
	conn   *grpc.ClientConn
}

type customCredential struct {
	apiKey string
}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"apikey": c.apiKey,
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func NewClient(server string, apiKey string, ops ...grpc.DialOption) *Client {
	ops = append(ops, grpc.WithPerRPCCredentials(customCredential{apiKey: apiKey}))
	conn, err := grpc.Dial(server, ops...)
	if err != nil {
		panic(err)
	}
	client := proto.NewMMApiClient(conn)
	return &Client{server: server, conn: conn, MMApiClient: client}
}

func (c *Client) Close() {
	c.conn.Close()
}

func NewUserClient(server string, ops ...grpc.DialOption) *UserClient {
	conn, err := grpc.Dial(server, ops...)
	if err != nil {
		panic(err)
	}
	client := proto.NewUserApiClient(conn)
	return &UserClient{server: server, conn: conn, UserApiClient: client}
}

func (c *UserClient) Close() {
	c.conn.Close()
}
