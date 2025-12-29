package providers

import (
	"context"
	"errors"
	"ezqueue/auth"

	"google.golang.org/api/idtoken"
)

type GoogleProvider struct {
	ClientID string
}

func (g *GoogleProvider) Name() string {
	return "google"
}

func (g *GoogleProvider) Verify(ctx context.Context, token string) (*auth.UserInfo, error) {
	payload, err := idtoken.Validate(ctx, token, g.ClientID)
	if err != nil {
		return nil, errors.New("invalid google token")
	}

	sub, _ := payload.Claims["sub"].(string)
	email, _ := payload.Claims["email"].(string)

	return &auth.UserInfo{
		Provider:       g.Name(),
		ProviderUserID: sub,
		Email:          email,
	}, nil
}
