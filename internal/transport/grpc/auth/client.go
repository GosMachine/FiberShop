package auth

import (
	"context"
	authv1 "github.com/GosMachine/protos/gen/go/auth"
)

func (c *Client) EmailVerified(ctx context.Context, email string) (*authv1.EmailVerifiedResponse, error) {
	resp, err := c.api.EmailVerified(ctx, &authv1.EmailVerifiedRequest{
		Email: email,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) Login(ctx context.Context, email, password, ip, rememberMe string) (string, error) {
	resp, err := c.api.Login(ctx, &authv1.LoginRequest{
		Email:      email,
		Password:   password,
		IP:         ip,
		RememberMe: rememberMe,
	})
	if err != nil {
		return "", err
	}
	return resp.Token, nil
}

func (c *Client) Register(ctx context.Context, email, password, ip, rememberMe string) (string, error) {
	resp, err := c.api.Register(ctx, &authv1.RegisterRequest{
		Email:      email,
		Password:   password,
		IP:         ip,
		RememberMe: rememberMe,
	})
	if err != nil {
		return "", err
	}

	return resp.Token, nil
}

func (c *Client) OAuth(ctx context.Context, email, ip string) (string, error) {
	resp, err := c.api.OAuth(ctx, &authv1.OAuthRequest{
		Email: email,
		IP:    ip,
	})
	if err != nil {
		return "", err
	}
	return resp.Token, nil
}
