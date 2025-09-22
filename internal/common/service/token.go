package service

import (
	"IdentifyService/internal/system/model"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type ITokenLogic interface {
	// 添加令牌到黑名单
	AddToBlacklist(ctx context.Context, jti string, operatorID string) (err error)
	// 判断令牌是否在黑名单中
	InBlacklist(ctx context.Context, jti string) (exists bool, err error)
}

type IToken interface {
	// 生成token
	Generate(ctx context.Context, data map[string]interface{}) (tokenStr string, err error)
	// 解析token: 只负责解析token，不负责验证token（如有效期等）
	Parse(ctx context.Context, tokenStr string) (customClaims *CustomClaims, err error)
	// 验证token(有效期、是否生效、是否加入了黑名单)
	Validate(ctx context.Context, customClaims *CustomClaims) (err error)
	// 令牌内省
	Introspect(ctx context.Context, tokenStr string) (user *model.IntrospectRes, err error)
	// 撤销令牌
	Revoke(ctx context.Context, tokenStr string, userID string) (err error)
}

var localTokenService IToken

func RegisterTokenService(logicsToken ITokenLogic) {
	localTokenService = &Token{
		ServerName: "IdentifyService",
		CacheKey:   "IdentifyService",
		Timeout:    3600 * 18,
		EncryptKey: []byte("IdentifyService1"),
		signer:     NewSigner("IdentifyService1"),

		logicsToken: logicsToken,
	}
}

func TokenService() IToken {
	if localTokenService == nil {
		panic("implement not found for interface IToken, forgot register?")
	}
	return localTokenService
}

var (
	ErrTokenInvalid               error = errors.New("无效的令牌")
	ErrTokenExpired               error = errors.New("令牌已过期")
	ErrTokenNotActive             error = errors.New("令牌未生效")
	ErrTokenUserInBlacklist       error = errors.New("账户被禁用")
	ErrTokenClientTypeInBlacklist error = errors.New("客户端被禁用")
)

type CustomClaims struct {
	CustomData map[string]interface{}
	jwt.RegisteredClaims
}

type Token struct {
	ServerName   string
	CacheKey     string
	Timeout      int64
	EncryptKey   []byte
	ExcludePaths g.SliceStr
	signer       *Signer

	logicsToken ITokenLogic
}

// data 为自定义数据
// user_id
// user_name
// user_nick_name
// org_id
// role_ids: []
// client_type
func (m *Token) Generate(ctx context.Context, data map[string]interface{}) (token string, err error) {
	data["client_type"] = "web"
	token, err = m.signer.Sign(CustomClaims{
		CustomData: data,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(m.Timeout))),
			ID:        uuid.New().String(),
		},
	})
	if err != nil {
		return
	}

	return
}

func (m *Token) Parse(ctx context.Context, tokenStr string) (customClaims *CustomClaims, err error) {
	customClaims, err = m.signer.Parse(tokenStr)
	if err != nil {
		return nil, err
	}

	return customClaims, nil
}

func (m *Token) Validate(ctx context.Context, customClaims *CustomClaims) (err error) {
	// 判断令牌是否过期
	if customClaims.ExpiresAt.Time.Before(time.Now()) {
		return ErrTokenExpired
	}
	// 判断令牌是否生效
	if customClaims.NotBefore.Time.After(time.Now()) {
		return ErrTokenNotActive
	}

	// 判断令牌是否在黑名单中
	exists, err := m.logicsToken.InBlacklist(ctx, customClaims.ID)
	if err != nil {
		return err
	}
	if exists {
		return ErrTokenExpired
	}

	return
}

func (m *Token) Introspect(ctx context.Context, token string) (user *model.IntrospectRes, err error) {
	customClaims, err := m.Parse(ctx, token)
	if err != nil {
		return nil, err
	}

	err = m.Validate(ctx, customClaims)
	if err != nil {
		return nil, err
	}

	userID, ok := customClaims.CustomData["user_id"].(string)
	if !ok {
		err = fmt.Errorf("%w: user_id not found", ErrTokenInvalid)
		return
	}
	userName, ok := customClaims.CustomData["user_name"].(string)
	if !ok {
		err = fmt.Errorf("%w: user_name not found", ErrTokenInvalid)
		return
	}
	userNickname, ok := customClaims.CustomData["user_nickname"].(string)
	if !ok {
		err = fmt.Errorf("%w: user_nickname not found", ErrTokenInvalid)
		return
	}
	orgID, ok := customClaims.CustomData["org_id"].(string)
	if !ok {
		err = fmt.Errorf("%w: org_id not found", ErrTokenInvalid)
		return
	}
	varRoleIDs, ok := customClaims.CustomData["role_ids"].([]interface{})
	if !ok {
		err = fmt.Errorf("%w: role_ids not found", ErrTokenInvalid)
		return
	}
	roleIDs := make([]int64, 0, len(varRoleIDs))
	for _, roleID := range varRoleIDs {
		v, ok := roleID.(float64)
		if !ok {
			err = fmt.Errorf("%w: role_ids invalid", ErrTokenInvalid)
			return
		}
		roleIDs = append(roleIDs, int64(v))
	}
	user = &model.IntrospectRes{
		UserID:       userID,
		UserName:     userName,
		UserNickname: userNickname,
		OrgID:        orgID,
		RoleIDs:      roleIDs,
	}
	return
}

func (m *Token) Revoke(ctx context.Context, token string, userID string) (err error) {
	jti, err := m.getJti(token)
	if err != nil {
		return err
	}

	err = m.logicsToken.AddToBlacklist(ctx, jti, userID)
	if err != nil {
		return err
	}

	return
}

func (m *Token) getJti(token string) (jti string, err error) {
	parts := strings.Split(token, ".")
	if len(parts) < 3 {
		return "", ErrTokenInvalid
	}

	data, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", ErrTokenInvalid
	}

	var i map[string]interface{}
	err = json.Unmarshal(data, &i)
	if err != nil {
		return "", ErrTokenInvalid
	}

	jti, ok := i["jti"].(string)
	if !ok {
		return "", ErrTokenInvalid
	}

	return jti, nil
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
	}, jwt.WithoutClaimsValidation())

	if err != nil {
		err = fmt.Errorf("%w: %v", ErrTokenInvalid, err)
		return
	}

	if !token.Valid {
		err = fmt.Errorf("%w: 令牌格式不正确", ErrTokenInvalid)
		return
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		err = fmt.Errorf("%w: 令牌数据格式不正确", ErrTokenInvalid)
		return
	}
	return
}
