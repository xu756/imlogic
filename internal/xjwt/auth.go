package xjwt

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/xu756/imlogic/internal/tool"
	"github.com/xu756/imlogic/internal/xerr"
)

type AuthInfo struct {
	ID     int64  `json:"Id"`     // 用户ID
	Role   int64  `json:"Role"`   // 用户角色
	Issuer string `json:"Issuer"` // 签发者
}

func (user *AuthInfo) GetStrId() string {
	return fmt.Sprintf("%d", user.ID)

}

type Jwt struct {
	SignKey string // 秘钥
	Expire  int64  // 过期时间
	Issuer  string // token颁发者
}

// NewJwt 初始化jwt
func NewJwt(j Jwt) *Jwt {
	return &Jwt{
		SignKey: j.SignKey,
		Expire:  j.Expire,
		Issuer:  j.Issuer,
	}
}

// token过期时间
func (j *Jwt) expireAtTime() time.Time {
	timezone := tool.TimeNowInTimeZone()
	expire := time.Duration(j.Expire) * time.Minute
	return timezone.Add(expire)
}

// customJwtClaims
//
//	 结构体实现了 Claims 接口继承了  Valid() 方法
//		jwt 规定了7个官方字段，提供使用:
//		  - iss (issuer)：发布者
//		  - sub (subject)：主题
//		  - iat (Issued At)：生成签名的时间
//		  - exp (expiration time)：签名过期时间
//		  - aud (audience)：观众，相当于接受者
//		  - nbf (Not Before)：生效时间
//		  - jti (jwt ID)：编号

type customJwtClaims struct {
	User AuthInfo `json:"user"`
	jwt.RegisteredClaims
}

// NewJwt 生成jwt，返回 token 字符串
func (j *Jwt) NewJwt(userId int64, role int64) (string, error) {
	c := customJwtClaims{
		User: AuthInfo{
			ID:     userId,
			Role:   role,
			Issuer: j.Issuer,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(tool.TimeNowInTimeZone()), // 生效时间
			IssuedAt:  jwt.NewNumericDate(tool.TimeNowInTimeZone()), // 签发时间
			ExpiresAt: jwt.NewNumericDate(j.expireAtTime()),         // 签名过期时间
		},
	}
	// 根据 claims 生成token对象
	token, err := j.createToken(c)
	if err != nil {
		return "", xerr.LogOut()
	}
	return token, nil
}

// createToken 生成token
func (j *Jwt) createToken(claims customJwtClaims) (string, error) {
	// 使用HS256算法进行token生成
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(j.SignKey))
}

// GetTokenFromHeader 获取请求头中的token
func (j *Jwt) GetTokenFromHeader(r *http.Request, headerName string) (*AuthInfo, error) {

	authHeader := r.Header.Get(headerName)
	if authHeader == "" {
		return nil, xerr.NewMsgError("请求头中auth为空")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, xerr.NewSystemError("请求头中auth格式有误")
	}
	token, err := j.parseTokenString(parts[1])
	if err != nil {
		validationErr, ok := err.(*jwt.ValidationError)
		if ok {
			if validationErr.Errors == jwt.ValidationErrorMalformed {
				return nil, xerr.NewMsgError("token 格式错误")
			} else if validationErr.Errors == jwt.ValidationErrorExpired {
				return nil, xerr.NewMsgError("token 过期")
			}
		}
		return nil, xerr.NewMsgError("token 无效")
	}
	if claims, ok := token.Claims.(*customJwtClaims); ok && token.Valid {
		return &claims.User, nil
	}
	return nil, xerr.NewMsgError("token 无效")

}

// parseTokenString 解析token
func (j *Jwt) parseTokenString(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &customJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SignKey), nil
	})
}
