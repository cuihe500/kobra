package jwt

import "time"

type JWT struct {
	Issuer string
	exp    time.Time
	jwt.StandardClaims
}
