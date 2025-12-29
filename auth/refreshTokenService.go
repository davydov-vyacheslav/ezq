package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"time"

	"cloud.google.com/go/firestore"
)

func GenerateRefreshToken() (string, time.Time) {
	b := make([]byte, 32)
	rand.Read(b)

	return base64.StdEncoding.EncodeToString(b),
		time.Now().Add(30 * 24 * time.Hour)
}

func HashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

type RefreshTokenRepo struct {
	Client *firestore.Client
}

type RefreshTokenModel struct {
	UserID    string    `firestore:"user_id"`
	ExpiresAt time.Time `firestore:"expires_at"`
	CreatedAt time.Time `firestore:"created_at"`
}

func (r *RefreshTokenRepo) Save(
	ctx context.Context,
	tokenHash string,
	userID string,
	expiresAt time.Time,
) error {
	_, err := r.Client.Collection("refresh_tokens").
		Doc(tokenHash).
		Set(ctx, RefreshTokenModel{
			UserID:    userID,
			ExpiresAt: expiresAt,
			CreatedAt: time.Now(),
		})

	return err
}

func (r *RefreshTokenRepo) Get(
	ctx context.Context,
	tokenHash string,
) (*RefreshTokenModel, error) {
	doc, err := r.Client.Collection("refresh_tokens").
		Doc(tokenHash).
		Get(ctx)
	if err != nil {
		return nil, err
	}

	var rt RefreshTokenModel
	doc.DataTo(&rt)
	return &rt, nil
}

func (r *RefreshTokenRepo) Delete(
	ctx context.Context,
	tokenHash string,
) error {
	_, err := r.Client.Collection("refresh_tokens").
		Doc(tokenHash).
		Delete(ctx)
	return err
}
