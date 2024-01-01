package auth

import (
	"context"
	"fmt"
	authv1 "github.com/GosMachine/protos/gen/go/auth"
)

func (c *Client) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	const op = "grpc.auth.IsAdmin"

	resp, err := c.api.IsAdmin(ctx, &authv1.IsAdminRequest{
		UserId: userID,
	})
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return resp.IsAdmin, nil
}

func (c *Client) IsUserLoggedIn(ctx context.Context, token string) (bool, string, error) {
	const op = "grpc.auth.IsUserLoggedIn"

	resp, err := c.api.IsUserLoggedIn(ctx, &authv1.IsUserLoggedInRequest{
		Token: token,
	})
	if err != nil {
		return false, "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.IsUserLoggedIn, resp.Token, nil
}

func (c *Client) Login(ctx context.Context, email, password, ip, rememberMe string) (string, error) {
	const op = "grpc.auth.Login"

	resp, err := c.api.Login(ctx, &authv1.LoginRequest{
		Email:      email,
		Password:   password,
		IP:         ip,
		RememberMe: rememberMe,
	})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.Token, nil
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

	return resp.Token, nil
}
