package xjwt

import (
	"imlogic/common/config"
	"imlogic/internal/tool"
	"imlogic/internal/xerr"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v5"
)

type AuthInfo struct {
	UserId   int64  `json:"id"`   // 用户ID
	UserUuid string `json:"uuid"` // 用户UUID
}

type Jwt struct {
	SignKey string // 秘钥
	Expire  int64  // 过期时间
}

func NewJwt() *Jwt {
	return &Jwt{
		SignKey: config.RunData().JwtConfig.SignKey,
		Expire:  config.RunData().JwtConfig.Expire,
	}
}

// token过期时间
func (j *Jwt) expireAtTime() time.Time {
	timezone := tool.TimeNow()
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

// NewJwtToken 生成jwt，返回 token 字符串
func (j *Jwt) NewJwtToken(UserId int64, uuid string) (string, error) {
	c := customJwtClaims{
		User: AuthInfo{
			UserId:   UserId,
			UserUuid: uuid,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(tool.TimeNow()),   // 生效时间
			IssuedAt:  jwt.NewNumericDate(tool.TimeNow()),   // 签发时间
			ExpiresAt: jwt.NewNumericDate(j.expireAtTime()), // 签名过期时间
		},
	}
	// 根据 claims 生成token对象
	token, err := j.createToken(c)
	if err != nil {
		return "", xerr.RoleErr("jwt生成失败")
	}
	return token, nil
}

// createToken 生成token
func (j *Jwt) createToken(claims customJwtClaims) (string, error) {
	// 使用HS256算法进行token生成
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(j.SignKey))
}

// GetUserInfoFromHeardToken 从token解析
func (j *Jwt) GetUserInfoFromHeardToken(c *app.RequestContext) (*AuthInfo, error) {
	token, err := j.parseTokenString(string(c.GetHeader("Authorization")))
	if err != nil {
		return nil, xerr.RoleErr("jwt解析失败")
	}
	if claims, ok := token.Claims.(*customJwtClaims); ok && token.Valid {
		return &claims.User, nil
	}
	return nil, xerr.RoleErr("jwt解析失败")
}

// GetUserInfoFromCookieToken 从token解析
func (j *Jwt) GetUserInfoFromCookieToken(c *app.RequestContext) (*AuthInfo, error) {
	token, err := j.parseTokenString(string(c.Cookie("ImlogicToken")))
	if err != nil {
		return nil, xerr.RoleErr("jwt解析失败")
	}
	if claims, ok := token.Claims.(*customJwtClaims); ok && token.Valid {
		return &claims.User, nil
	}
	return nil, xerr.RoleErr("jwt解析失败")
}

// parseTokenString 解析token
func (j *Jwt) parseTokenString(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &customJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SignKey), nil
	})
}
