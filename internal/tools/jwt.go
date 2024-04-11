package tools

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/common"
	"log/slog"
	"time"
)

type Token struct {
	UUID string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateNewJwtToken(uuid uuid.UUID, username string) (string, error) {
	t, err1 := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		Token{
			UUID: uuid.String(),
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "kobra",
				Subject:   "users",
				Audience:  []string{username},
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
			},
		},
	).SignedString([]byte(common.JWTSecret))
	if err1 != nil {
		slog.Error("签发Jwt Token令牌失败！", "reason", err1)
		return "", errors.New("签发Jwt Token令牌失败！")
	}
	return t, nil
}
func ValidateAndParseJwtToken(tokenString string) (*Token, error) {
	tk, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.JWTSecret), nil
	},
		jwt.WithLeeway(time.Second*5),
		jwt.WithIssuer("kobra"),
		jwt.WithSubject("users"),
	)
	if err != nil {
		slog.Error("验证token错误！", "reason", err)
		return nil, errors.New("验证token错误！")
	}
	if !tk.Valid {
		slog.Error("该Token未验证！")
		return nil, errors.New("该Token未验证！")
	}
	if token, ok := tk.Claims.(*Token); !ok {
		slog.Error("该Token转换错误！", "reason", err)
		return nil, errors.New("转换Token错误！")
	} else {
		return token, nil
	}
}
