package auth

import "context"

type UserInfo struct {
	Provider       string
	ProviderUserID string
	Email          string
}

type Provider interface {
	Name() string
	Verify(ctx context.Context, token string) (*UserInfo, error)
}
