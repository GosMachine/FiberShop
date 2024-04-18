package auth

import (
	"context"
	authv1 "github.com/GosMachine/protos/gen/go/auth"
)

func (c *Client) EmailVerified(ctx context.Context, email string) (bool, error) {
	resp, err := c.api.EmailVerified(ctx, &authv1.EmailVerifiedRequest{
		Email: email,
	})
	if err != nil {
		return false, err
	}
	return resp.EmailVerified, nil
}

func (c *Client) EmailVerify(ctx context.Context, email string) error {
	_, err := c.api.EmailVerify(ctx, &authv1.EmailVerifyRequest{
		Email: email,
	})
	return err
}

func (c *Client) ChangePass(ctx context.Context, email, password, ip string) (string, error) {
	resp, err := c.api.ChangePass(ctx, &authv1.ChangePassRequest{
		Email:    email,
		Password: password,
		IP:       ip,
	})
	if err != nil {
		return "", err
	}
	return resp.Token, nil
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
