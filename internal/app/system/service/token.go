package service

import (
	"IdentifyService/internal/app/system/model"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type IToken interface {
	// 生成token
	GenerateToken(ctx context.Context, key string, data interface{}) (keys string, err error)
	// 解析token
	ParseToken(r *ghttp.Request) (*CustomClaims, error)
	// 令牌内省
	Introspect(ctx context.Context, token string) (user *model.IntrospectRes, err error)
}

var localTokenService IToken

func RegisterTokenService() {
	localTokenService = &Token{
		ServerName: "IdentifyService",
		CacheKey:   "IdentifyService",
		Timeout:    3600 * 18,
		EncryptKey: []byte("IdentifyService1"),
		signer:     NewSigner("IdentifyService1"),
	}
}

func TokenService() IToken {
	if localTokenService == nil {
		panic("implement not found for interface IToken, forgot register?")
	}
	return localTokenService
}

const (
	ErrTokenInvalid string = "无效的token"
	ErrTokenExpired string = "token已过期"
)

type CustomClaims struct {
	Data interface{}
	jwt.RegisteredClaims
}

type Token struct {
	ServerName   string
	CacheKey     string
	Timeout      int64
	EncryptKey   []byte
	ExcludePaths g.SliceStr
	signer       *Signer
}

func (m *Token) GenerateToken(ctx context.Context, key string, data interface{}) (token string, err error) {
	if len(key) < 32 {
		err = gerror.New("key length must more than 32")
		return
	}

	tokenStr, err := m.signer.Sign(CustomClaims{
		data,
		jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(m.Timeout))),
			ID:        uuid.New().String(),
		},
	})
	if err != nil {
		return
	}
	token, err = m.encrypt(tokenStr)
	if err != nil {
		return
	}

	return
}

func (m *Token) ParseToken(r *ghttp.Request) (claims *CustomClaims, err error) {
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		err = errors.New("bearer token is required")
		return
	}
	tokenStr := strings.TrimPrefix(authorization, "Bearer ")
	if tokenStr == "" {
		err = errors.New("bearer token is required")
		return
	}

	token, err := m.decrypt(tokenStr)
	if err != nil {
		return nil, err
	}

	customClaims, err := m.signer.Parse(token)
	if err != nil {
		return nil, err
	}
	return customClaims, nil
}

func (m *Token) Introspect(ctx context.Context, token string) (user *model.IntrospectRes, err error) {
	return
}

func (m *Token) encrypt(plainText string) (encryptStr string, err error) {
	if plainText == "" {
		err = gerror.New("encrypt plainText can not be empty")
		return
	}

	token, err := gaes.Encrypt([]byte(plainText), m.EncryptKey)
	if err != nil {
		err = fmt.Errorf("encrypt error: %w", err)
		return
	}
	encryptStr = gbase64.EncodeToString(token)
	return
}

func (m *Token) decrypt(cipherText string) (decryptStr string, err error) {
	if cipherText == "" {
		err = gerror.New("cipherText can not be empty")
		return
	}
	token64, err := gbase64.Decode([]byte(cipherText))
	if err != nil {
		err = gerror.New("decode error")
		return
	}
	token, err := gaes.Decrypt(token64, m.EncryptKey)
	if err != nil {
		err = gerror.New("decrypt error")
		return
	}
	decryptStr = string(token)
	return
}

type Signer struct {
	SigningKey []byte
}

func NewSigner(JwtTokenSignKey string) *Signer {
	return &Signer{
		[]byte(JwtTokenSignKey),
	}
}

func (j *Signer) Sign(claims CustomClaims) (string, error) {
	tokenPartA := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenPartA.SignedString(j.SigningKey)
}

func (j *Signer) Parse(tokenString string) (claims *CustomClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return
	}
	if token == nil {
		err = errors.New(ErrTokenInvalid)
		return
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		err = errors.New(ErrTokenInvalid)
		return
	}
	return
}
