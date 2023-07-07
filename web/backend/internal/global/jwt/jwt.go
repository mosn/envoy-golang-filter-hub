package jwt

import (
	"envoy-golang-filter-hub/config"
	"envoy-golang-filter-hub/utils"
	"github.com/golang-jwt/jwt"
)

type Payload struct {
	GitHubUsername string `json:"github_username"`
	GitHubID       int64  `json:"github_id"`
	AvatarURL      string `json:"avatar_url"`
}

type Claims struct {
	Payload
	jwt.StandardClaims
}

// CreateToken 签发用户Token
func CreateToken(payload Payload) (string, error) {
	claims := Claims{
		Payload: payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: utils.GetUnix() + config.Get().JWT.AccessExpire,
			IssuedAt:  utils.GetUnix(),
			Issuer:    config.Get().JWT.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(config.Get().JWT.AccessSecret))
	return token, err
}

// ParseToken 解析用户Token
func ParseToken(token string) (claims *Claims, ok bool) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Get().JWT.AccessSecret), nil
		},
	)
	if err != nil || tokenClaims == nil {
		//logx.Error("Failed to parse token: %v", err)
		return nil, false
	}
	if claims, ok = tokenClaims.Claims.(*Claims); !ok || !tokenClaims.Valid {
		return nil, false
	}

	return claims, true
}
