package models

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"time"
)

type Session struct {
	UserID     string
	PlainText  string
	Hash       []byte
	Disabled   bool
	ExpiresAt  time.Time
	CreatedAt  time.Time
	DisabledAt time.Time
}

func CreateSession(userID string, ttl time.Duration) (*Session, error) {
	session := &Session{
		UserID:    userID,
		ExpiresAt: time.Now().Add(ttl),
		CreatedAt: time.Now(),
	}
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	session.PlainText = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := sha256.Sum256([]byte(session.PlainText))
	session.Hash = hash[:]
	return session, nil
}
