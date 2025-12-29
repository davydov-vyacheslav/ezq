package auth

import (
	"context"
	"ezqueue/models"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
)

type UserRepo struct {
	Client *firestore.Client
}

func (r *UserRepo) FindOrCreateUser(
	ctx context.Context,
	info UserInfo,
) (string, error) {

	identityID := fmt.Sprintf("%s:%s", info.Provider, info.ProviderUserID)
	identityRef := r.Client.Collection("user_identities").Doc(identityID)

	// 1️⃣ Пробуем найти identity
	doc, err := identityRef.Get(ctx)
	if err == nil {
		// identity существует → берём user_id
		userID, _ := doc.Data()["user_id"].(string)
		return userID, nil
	}

	// 2️⃣ Identity нет → создаём нового user
	userID := uuid.NewString()
	userRef := r.Client.Collection("users").Doc(userID)

	batch := r.Client.Batch()

	// users
	batch.Set(userRef, models.User{
		Email:     info.Email,
		CreatedAt: time.Now(),
	})

	// user_identities
	batch.Set(identityRef, models.UserIdentity{
		UserID:         userID,
		Provider:       info.Provider,
		ProviderUserID: info.ProviderUserID,
		Email:          info.Email,
		CreatedAt:      time.Now(),
	})

	_, err = batch.Commit(ctx)
	if err != nil {
		return "", err
	}

	return userID, nil
}
