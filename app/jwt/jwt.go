package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	JWT_SECRET     = "112d8904aswd9adqwd19asdas1d9qw0d74asd1a"
	JWT_TOKEN_LIFE = 2592000
	Key            = "token"
)

type Claims struct {
	MemberId   string
	MemberName string
	jwt.StandardClaims
}

func GenerateToken(memberId, memberName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(JWT_TOKEN_LIFE) * time.Second)
	claims := &Claims{
		memberId,
		memberName,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(JWT_SECRET))
	return token, err
}

func ParseToken(token string) (memberId, memberName string, err error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		return "", "", err
	}

	claims, ok := tokenClaims.Claims.(*Claims)
	if !ok || !tokenClaims.Valid {
		return "", "", errors.New("tokenClaims invalid")
	}
	return claims.MemberId, claims.MemberName, nil
}
