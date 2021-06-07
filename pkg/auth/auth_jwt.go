package auth

import (
	"go-skeleton/pkg/common"
	"go-skeleton/pkg/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang-module/carbon"
)

var JwtKey = []byte(config.Conf.AppConfig.JwtKey)

type MyClaims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

//生成jwt token 有效期10分钟
func GenerateToken(id int) (string, error) {
	expireTime := carbon.Now().AddMinutes(10).ToTimestamp()

	SetClaims := MyClaims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "gin_test",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

// 验证token
func CheckToken(token string) (*MyClaims, *common.CodeError) {
	var claims MyClaims

	claimsToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				//token错误
				return nil, common.TokenWrongError
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				//token过期
				return nil, common.TokenRuntimeError
			} else {
				//token格式错误
				return nil, common.TokenTypeWrongError
			}
		}
	}

	if claimsToken == nil {
		return nil, common.TokenWrongError
	}

	if claim, ok := claimsToken.Claims.(*MyClaims); ok && claimsToken.Valid {
		return claim, nil
	} else {
		return nil, common.TokenWrongError
	}
}
