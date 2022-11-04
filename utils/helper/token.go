package helper

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"time"
)

var (
	hmacSampleSecret = viper.GetString("jwt.hmacSampleSecret")
	issuer           = viper.GetString("jwt.issuer")
)

type CustomerUserInfo struct {
	UserName string
	UserId   uint
}

type CustomClaimsExample struct {
	*jwt.StandardClaims
	CustomerUserInfo
}

// GenerateToken
//
//	@Description: token生成
func GenerateToken(userInfo CustomerUserInfo) (string, error) {
	claims := CustomClaimsExample{
		CustomerUserInfo: CustomerUserInfo{
			UserName: userInfo.UserName,
			UserId:   userInfo.UserId,
		},
		StandardClaims: &jwt.StandardClaims{
			Issuer:    issuer,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24 * 5).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func ParseToken(tokenString string) (*CustomerUserInfo, error) {
	fmt.Printf(tokenString)
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsExample{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(hmacSampleSecret), nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaimsExample)
	if !ok {
		fmt.Println("ok")
		return nil, errors.New("token错误")
	}
	return &claims.CustomerUserInfo, nil
}
