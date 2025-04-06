package jwt

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var secretKey = []byte(`139db9c9a32bd7ad433dd4717455cb11ae9e631e15081c8306ee31cc84ca899886c3e17fd7590bf1492881e97eeae0dfb12107bb8cfa89144a2e788cdc42359b89eed83a1712d79009357475034931897eafa80aaafc233ddf2846a2c9a9aad1e6694bb409f65a5c29738953116580afba156c0e5105629173c1e8b6bdf595c9b8c6a776435d1def02365fdcfbe580790ca7c70f1a72e40dc0147dfc7299cd3852f596c13864061a537c322c148b73774d45ef3007b5027e81832de9c65534f66f0cdd8da91543c35f7fc3ad71cb86ea26214904f171cceaaacd3800d9b95fe61ee83787d3661819775ccb092af8c54c70e8d7b239c70e13caab718f8e1a7d16`)
var authDuration = time.Hour * 24 * 30

type JWT struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

func GetClaim(tokenStr string) (string, time.Time, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", time.Time{}, err
	}
	if !token.Valid {
		return "", time.Time{}, errors.New("unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", time.Time{}, errors.New("invalid claims")
	}

	sub, exists := claims["sub"].(string)
	if !exists {
		return "", time.Time{}, errors.New("subject not found")
	}

	expiresAt, ok := claims["exp"].(float64)
	if !ok {
		return "", time.Time{}, errors.New("expiration time not found")
	}

	return sub, time.Unix(int64(expiresAt), 0), nil
}

func GenerateToken(userID uint) JWT {
	accessToken, _ := accessToken(userID)
	refreshToken := refreshToken(accessToken)

	_, expiresAt, _ := GetClaim(accessToken)

	return JWT{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	}
}

func accessToken(userID uint) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.StandardClaims{
			Subject:   fmt.Sprintf("%v", userID),
			ExpiresAt: now.Add(authDuration).Unix(),
			NotBefore: now.Unix(),
			IssuedAt:  now.Unix(),
			Id:        uuid.New().String(),
		})

	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func refreshToken(accessToken string) string {
	t := uuid.NewSHA1(uuid.Must(uuid.NewRandom()), []byte(accessToken)).String()
	refresh := base64.URLEncoding.EncodeToString([]byte(t))
	return strings.ToUpper(strings.TrimRight(refresh, "="))
}
