package libSecurity

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
)

const idCardEncryptKeyPath = "server.security.idCardEncryptKey"

// EncryptIDCard 使用对称加密对身份证号进行加密，返回Base64字符串
func EncryptIDCard(idCard string) (string, error) {
	if idCard == "" {
		return "", nil
	}
	key, err := loadAESKey()
	if err != nil {
		return "", err
	}
	return EncryptAES(idCard, key)
}

// DecryptIDCard 解密身份证号
func DecryptIDCard(cipherText string) (string, error) {
	if cipherText == "" {
		return "", nil
	}
	key, err := loadAESKey()
	if err != nil {
		return "", err
	}
	return DecryptAES(cipherText, key)
}

// EncryptAES 使用AES-GCM进行加密
func EncryptAES(plaintext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	cipherBytes := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(cipherBytes), nil
}

// DecryptAES 使用AES-GCM进行解密
func DecryptAES(cipherText string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	if len(data) < gcm.NonceSize() {
		return "", errors.New("cipherText too short")
	}
	nonce := data[:gcm.NonceSize()]
	cipherBytes := data[gcm.NonceSize():]
	plainBytes, err := gcm.Open(nil, nonce, cipherBytes, nil)
	if err != nil {
		return "", err
	}
	return string(plainBytes), nil
}

// MaskIDCard 身份证号脱敏，保留前3后4
func MaskIDCard(idCard string) string {
	if idCard == "" {
		return ""
	}
	runes := []rune(idCard)
	length := len(runes)
	if length <= 7 {
		return "***"
	}
	for i := 3; i < length-4; i++ {
		runes[i] = '*'
	}
	return string(runes)
}

// MaskMobile 手机号脱敏，保留前3后4
func MaskMobile(mobile string) string {
	if mobile == "" {
		return ""
	}
	runes := []rune(mobile)
	length := len(runes)
	if length <= 7 {
		return "***"
	}
	for i := 3; i < length-4; i++ {
		runes[i] = '*'
	}
	return string(runes)
}

// MaskRealName 对真实姓名进行脱敏，保留首个字符
func MaskRealName(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return ""
	}
	runes := []rune(name)
	if len(runes) == 1 {
		return string(runes[:1]) + "*"
	}
	masked := make([]rune, len(runes))
	masked[0] = runes[0]
	for i := 1; i < len(runes); i++ {
		masked[i] = '*'
	}
	return string(masked)
}

func loadAESKey() ([]byte, error) {
	keyStr := g.Cfg().MustGet(context.Background(), idCardEncryptKeyPath).String()
	keyBytes := []byte(keyStr)
	if len(keyBytes) != 32 {
		return nil, errors.New("idCardEncryptKey must be 32 bytes")
	}
	return keyBytes, nil
}
