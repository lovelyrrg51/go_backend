package dtos

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTStandardClaim struct {
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}
