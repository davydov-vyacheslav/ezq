package models

import "time"

type User struct {
	ID    string `firestore:"-" json:"id"`
	Email string `firestore:"email" json:"email"` // TODO: unique
	//DisplayName string    `firestore:"displayName" json:"displayName"`
	//PhotoURL    string    `firestore:"photoURL" json:"photoURL"`
	CreatedAt time.Time `firestore:"createdAt" json:"createdAt"`
	//	FCMToken    string    `firestore:"fcmToken" json:"fcmToken"`
}

type UserIdentity struct {
	UserID         string    `firestore:"user_id"`
	Provider       string    `firestore:"provider"`
	ProviderUserID string    `firestore:"provider_user_id"`
	Email          string    `firestore:"email"`
	CreatedAt      time.Time `firestore:"created_at"`
}
