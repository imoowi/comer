package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
	Username string `json:"username"`
	UserId   uint   `json:"user_id"`
	RoleIds  string `json:"role_id"`
	jwt.StandardClaims
}

func GenToken(username string, userId uint, roleIds string) (string, error) {
	settingSecret := global.Config.GetString("jwt.secret")
	var MySecret = []byte(settingSecret)
	TokenExpireDuration := global.Config.GetDuration("jwt.timeout")
	// TokenExpireDuration := time.Hour * 2
	fmt.Println(`TokenExpireDuration is:`, TokenExpireDuration)
	// 创建一个我们自己的声明
	c := TokenClaims{
		username,
		userId,
		roleId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "lynkros-admin",                            // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

func ParseToken(tokenString string) (*TokenClaims, error) {
	var MySecret = []byte(global.Config.GetString("jwt.secret"))
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token.")
}
