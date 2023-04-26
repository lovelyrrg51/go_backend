package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/lovelyrrg51/go_backend/app/common"
	"github.com/lovelyrrg51/go_backend/app/config"
	"github.com/lovelyrrg51/go_backend/app/dtos"
	"github.com/lovelyrrg51/go_backend/app/logger"
)

func GenerateStandardJWT(id uuid.UUID) (*jwt.Token, *string, *common.AppError) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &dtos.JWTStandardClaim{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "GolangBackend",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Cfg.JWTSecretKey))
	if err != nil {
		logger.Error("Error when signed string token " + err.Error())
		return nil, nil, common.NewUnexpectedError("Unexpected error when signed string token " + err.Error())
	}

	return token, &tokenString, nil
}

func VerifyStandardJWTToken(signedToken string) (*uuid.UUID, *common.AppError) {
	token, err := jwt.ParseWithClaims(signedToken, &dtos.JWTStandardClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.JWTSecretKey), nil
	}, jwt.WithLeeway(5*time.Second))

	if err != nil {
		logger.Error("Error when parsed string token " + err.Error())
		return nil, common.NewUnexpectedError("Unexpected error when parsed string token " + err.Error())
	}

	claims, ok := token.Claims.(*dtos.JWTStandardClaim)
	if !ok || !token.Valid {
		logger.Error("Error when parsed string token " + err.Error())
		return nil, common.NewUnexpectedError("Unexpected error when signed string token " + err.Error())
	}

	return &claims.ID, nil
}
