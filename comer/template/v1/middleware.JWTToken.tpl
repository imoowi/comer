package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/spf13/viper"
)

type MyClaims struct {
	Username string `json:"username"`
	AdminId  string `json:"admin_id"`
	RoleId   string `json:"role_id"`
	jwt.StandardClaims
}

type WxUser struct {
	WxUserId string `json:"wx_user_id"`
	jwt.StandardClaims
}

func GenToken(username string, adminId string, roleId string) (string, error) {
	settingSecret := viper.GetString("settings.jwt.secret")
	fmt.Println(`settingSecret is:`, settingSecret)
	var MySecret = []byte(settingSecret)
	fmt.Println(`MySecret is:`, MySecret)
	TokenExpireDuration := viper.GetDuration("settings.jwt.timeout")
	// TokenExpireDuration := time.Hour * 2
	fmt.Println(`TokenExpireDuration is:`, TokenExpireDuration)
	// 创建一个我们自己的声明
	c := MyClaims{
		username,
		adminId,
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

func GenWxUserToken(wxUserid string) (string, error) {
	settingSecret := viper.GetString("settings.jwt.secret")
	fmt.Println(`settingSecret is:`, settingSecret)
	var MySecret = []byte(settingSecret)
	fmt.Println(`MySecret is:`, MySecret)
	TokenExpireDuration := viper.GetDuration("settings.jwt.timeout")
	// TokenExpireDuration := time.Hour * 2
	fmt.Println(`TokenExpireDuration is:`, TokenExpireDuration)
	// 创建一个我们自己的声明
	c := WxUser{
		wxUserid,
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

func ParseToken(tokenString string) (*MyClaims, error) {
	var MySecret = []byte(viper.GetString("settings.jwt.secret"))
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token.")
}

func ParseWxUserToken(tokenString string) (*WxUser, error) {
	var MySecret = []byte(viper.GetString("settings.jwt.secret"))
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &WxUser{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*WxUser); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token.")
}
