package app

import (
	"github.com/dgrijalva/jwt-go"
	"myproject/global"
	"time"
)

type Claims struct {
	AppKey string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(appKey,appSecret string) (string,error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey: appKey,
		AppSecret: appSecret,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: global.JWTSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	//根据传入的secret签名生成签名字符串
	token,err := tokenClaims.SignedString(GetJWTSecret())
	return token,err
}

//校验token
func ParseToken(token string) (*Claims,error) {
	//ParseWithClaims方法内部有具体的解码和校验过程
	tokenClaims,err := jwt.ParseWithClaims(token,&Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(),nil
	})
	if tokenClaims != nil{
		claims,ok := tokenClaims.Claims.(*Claims)
		//验证是否获得签名及签名是否过期
		if ok &&tokenClaims.Valid{
			return claims,nil
		}
	}
	return nil, err
}
