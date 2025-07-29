package JWT

import (
	"HQ/models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)
	
type LoginClaims struct {
	Username             string `json:"username"`
	Role                 int8   `json:"role"`
	UserId               int64  `json:"user_id"`
	jwt.RegisteredClaims        //内嵌标准声明
}

// customSecret 自定义密钥
var customSecret = []byte("HQgogogo")

// GenLoginToken 生成登录token
func GenLoginToken(loginParam models.LoginParam, ro int8, id int64, dt time.Duration) (string, error) {
	claims := LoginClaims{
		Username: loginParam.Username,
		Role:     ro,
		UserId:   id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(dt)),
			Issuer:    "HQ",
		},
	}
	// 适用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(customSecret)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*LoginClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &LoginClaims{}, func(token *jwt.Token) (i any, err error) {
		return customSecret, nil
	})
	if err != nil {
		return nil, err
	}
	//对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*LoginClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
