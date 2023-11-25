package xjwt

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v5"
	"github.com/xu756/imlogic/common/config"
	"github.com/xu756/imlogic/internal/tool"
	"github.com/xu756/imlogic/internal/xerr"
	"time"
)

type AuthInfo struct {
	UserID string `json:"UserID"` // 用户ID
	Role   string `json:"Role"`   // 用户角色
}

type Jwt struct {
	SignKey string // 秘钥
	Expire  int64  // 过期时间
}

func InitJwt() *Jwt {
	return &Jwt{
		SignKey: config.RunData.JwtConfig.SignKey,
		Expire:  config.RunData.JwtConfig.Expire,
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
func (j *Jwt) NewJwtToken(userId, role string) (string, error) {
	c := customJwtClaims{
		User: AuthInfo{
			UserID: userId,
			Role:   role,
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
		return "", xerr.ErrMsg(xerr.JwtCreateErr)
	}
	return token, nil
}

// createToken 生成token
func (j *Jwt) createToken(claims customJwtClaims) (string, error) {
	// 使用HS256算法进行token生成
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(j.SignKey))
}

// GetUserInfoFromToken 从token解析
func (j *Jwt) GetUserInfoFromToken(c *app.RequestContext) (*AuthInfo, error) {
	token, err := j.parseTokenString(string(c.Cookie("token")))
	if err != nil {
		return nil, xerr.ErrMsg(xerr.JwtParseErr)
	}
	if claims, ok := token.Claims.(*customJwtClaims); ok && token.Valid {
		return &claims.User, nil
	}
	return nil, xerr.ErrMsg(xerr.JwtParseErr)

}

// parseTokenString 解析token
func (j *Jwt) parseTokenString(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &customJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SignKey), nil
	})
}
