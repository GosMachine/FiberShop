package product

import (
	"context"
	productv1 "github.com/GosMachine/protos/gen/go/product"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Client struct {
	api productv1.ProductClient
}

func New(ctx context.Context, addr string, timeout int, retriesCount int) (*Client, error) {
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
		return nil, err
	}

	grpcClient := productv1.NewProductClient(cc)

	return &Client{
		api: grpcClient,
	}, nil
}
