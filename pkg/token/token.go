package token

import (
	"bug_busters/pkg/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func ValidateToken(tokenstr string) (bool, error) {
	_, err := ExtractClaims(tokenstr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractClaims(tokenstr string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenstr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		// Token imzosi HMAC bo'lishi kerak
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Load().SIGNING_KEY), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token: %s", tokenstr)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to parse token claims")
	}

	return claims, nil
}

// Token yaratish
func GenerateJWT(id, role, password string) (string, error) {
	accesstoken := jwt.New(jwt.SigningMethodHS256)

	// Access Token uchun claimlar
	accessClaim := accesstoken.Claims.(jwt.MapClaims)
	accessClaim["user_id"] = id
	accessClaim["role"] = role
	accessClaim["password"] = password
	accessClaim["iat"] = time.Now().Unix()
	accessClaim["exp"] = time.Now().Add(24 * time.Hour).Unix()

	con := config.Load()
	access, err := accesstoken.SignedString([]byte(con.SIGNING_KEY))
	if err != nil {
		return "", fmt.Errorf("error generating access token: %w", err)
	}

	return access, nil
}