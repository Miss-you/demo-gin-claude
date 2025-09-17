package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT Claims结构
type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

// JWTHelper JWT辅助工具
type JWTHelper struct {
	Secret []byte
	Issuer string
	Expire time.Duration
}

// NewJWTHelper 创建JWT辅助工具
func NewJWTHelper(secret string, expireHours int) *JWTHelper {
	return &JWTHelper{
		Secret: []byte(secret),
		Issuer: "demo-gin-test",
		Expire: time.Duration(expireHours) * time.Hour,
	}
}

// GenerateTestToken 生成测试Token
func (j *JWTHelper) GenerateTestToken(userID int, username, email string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.Expire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.Issuer,
			Subject:   fmt.Sprintf("%d", userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.Secret)
}

// GenerateExpiredToken 生成过期的Token
func (j *JWTHelper) GenerateExpiredToken(userID int, username, email string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
			Issuer:    j.Issuer,
			Subject:   fmt.Sprintf("%d", userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.Secret)
}

// ValidateTestToken 验证Token
func (j *JWTHelper) ValidateTestToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.Secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// ExtractClaims 从Token中提取Claims
func (j *JWTHelper) ExtractClaims(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.Secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("failed to extract claims")
}

// GetTokenWithInvalidSignature 生成签名错误的Token
func (j *JWTHelper) GetTokenWithInvalidSignature(userID int, username, email string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.Expire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.Issuer,
			Subject:   fmt.Sprintf("%d", userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用错误的密钥签名
	return token.SignedString([]byte("wrong-secret"))
}

// QuickToken 快速生成测试Token（使用默认配置）
func QuickToken(userID int) string {
	helper := NewJWTHelper("test_jwt_secret_key_for_testing_only", 24)
	token, _ := helper.GenerateTestToken(userID, fmt.Sprintf("user%d", userID), fmt.Sprintf("user%d@test.com", userID))
	return token
}