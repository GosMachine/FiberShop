package auth

import (
	"context"
	"fmt"
	authv1 "github.com/GosMachine/protos/gen/go/auth"
)

func (c *Client) User(ctx context.Context, email string) (*authv1.UserResponse, error) {
	const op = "grpc.auth.IsUserLoggedIn"

	resp, err := c.api.User(ctx, &authv1.UserRequest{
		Email: email,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}

func (c *Client) Login(ctx context.Context, email, password, ip, rememberMe, AuthMethod string) (string, error) {
	const op = "grpc.auth.Login"

	resp, err := c.api.Login(ctx, &authv1.LoginRequest{
		Email:      email,
		Password:   password,
		IP:         ip,
		RememberMe: rememberMe,
		AuthMethod: AuthMethod,
	})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return resp.GetToken(), nil
}

func (c *Client) Register(ctx context.Context, email, password, ip, rememberMe string) (string, error) {
	const op = "grpc.auth.Register"

	resp, err := c.api.Register(ctx, &authv1.RegisterRequest{
		Email:      email,
		Password:   password,
		IP:         ip,
		RememberMe: rememberMe,
	})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetToken(), nil
}
