package auth

import (
	"github.com/tiwariayush700/user-management/models"
)

type Service interface {
	GenerateUserToken(userID uint, role string) (string, error)
	AuthenticateUser(jwtTokenString string) (*models.UserLoginJWTClaims, error)
}
