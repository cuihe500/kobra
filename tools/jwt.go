package tools

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/common"
	"log/slog"
	"time"
)

func GenerateNewJwtToken(uuid uuid.UUID, username string) (string, error) {
	t, err1 := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":  "kobra",
			"exp":  time.Now().Add(time.Hour * 3).Unix(),
			"iat":  time.Now().Unix(),
			"nbf":  time.Now().Unix(),
			"sub":  "Users",
			"aud":  username,
			"uuid": uuid.String(),
		},
	).SignedString([]byte(common.JWTSecret))
	if err1 != nil {
		slog.Error("签发Jwt Token令牌失败！", "reason", err1)
		return "", errors.New("签发Jwt Token令牌失败！")
	}
	return t, nil
}
