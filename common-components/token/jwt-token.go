package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"time"
)

type JWTMaker struct {
	mySigningKey []byte
}

func NewJWTMaker(key string) Maker {
	return &JWTMaker{
		mySigningKey: []byte(key),
	}
}

type JWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	UserId      int64    `json:"id"`
	Username    string   `json:"name"`
	Permissions []string `json:"permissions"`
}

//var mySigningKey = []byte("cy&july")

// CreateToken Generate Token
func (maker *JWTMaker) CreateToken(id int64, userName string, expireDuration time.Duration) (string, error) {
	myClaims := &JWTClaims{
		Username:    userName,
		UserId:      id,
		Permissions: []string{},
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    "july",
			Subject:   "myToken",
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
		},
	}
	myToken := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	myTokenStr, err := myToken.SignedString(maker.mySigningKey)
	if err != nil {
		return "", err
	}

	return myTokenStr, nil
}

// VerifyToken Verify Token
func (maker *JWTMaker) VerifyToken(myToken string) (string, int64, bool) {
	if myToken == "" {
		return "", 0, false
	}

	tok, err := jwt.ParseWithClaims(myToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return maker.mySigningKey, nil
	})
	if err != nil {
		zap.L().Error("ParseWithClaims error", zap.Error(err))
		return "", 0, false
	}

	if claims, ok := tok.Claims.(*JWTClaims); ok && tok.Valid {
		return claims.Username, claims.UserId, true
	} else {
		zap.L().Error("", zap.Error(err))
		return "", 0, false
	}

}

// ParseUserIdByToken Parse UserId By Token
func (maker *JWTMaker) ParseUserIdByToken(tokenString string) (userId int64, err error) {
	if tokenString == "" {
		return 0, errors.New("token is nil")
	}
	// 使用密钥解码 JWT token
	tok, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return maker.mySigningKey, nil
	})
	if err != nil {
		zap.L().Error("ParseWithClaims error", zap.Error(err))
		return 0, err
	}
	// 检查 token 是否有效
	if claims, ok := tok.Claims.(*JWTClaims); ok && tok.Valid {
		return claims.UserId, nil
	} else {
		zap.L().Error("", zap.Error(err))
		return 0, err
	}
}
