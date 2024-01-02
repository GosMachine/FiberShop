package auth

import (
	"context"
	"fmt"
	authv1 "github.com/GosMachine/protos/gen/go/auth"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Client struct {
	api authv1.AuthClient
}

func New(
	ctx context.Context,
	addr string,
	timeout int,
	retriesCount int,
) (*Client, error) {
	const op = "grpc.auth.New"
	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(time.Duration(timeout) * time.Second),
	}
	cc, err := grpc.DialContext(ctx, addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			grpcretry.UnaryClientInterceptor(retryOpts...),
		))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	grpcClient := authv1.NewAuthClient(cc)

	return &Client{
		api: grpcClient,
	}, nil
}
