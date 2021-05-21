package services

import (
	"context"
	"fmt"
	"gin-test/pkg/gredis"
	"strings"
	"time"

	uuid "github.com/iris-contrib/go.uuid"
)

var UserTokenService = newUserTokenService()

func newUserTokenService() *userTokenService {
	return &userTokenService{}
}

//用户的refresh token
type userTokenService struct {
}

// 生成refreshToken
func (s *userTokenService) GenerateRefreshToken(userId int64) (string, error) {
	token := s.GetRefreshTokenByUserId(userId)
	if token == "" {
		uuidStr, _ := uuid.NewV4()
		token = strings.ReplaceAll(uuidStr.String(), "-", "")
		//30天过期时间
		tokenExpireDays := time.Hour * 24 * 40
		key := fmt.Sprintf("user:token2uid:%s", token)
		_ = gredis.GetRedis().Set(context.TODO(), key, userId, tokenExpireDays).Err()

		_ = s.SetUserIdToToken(userId, token)
	}
	return token, nil
}

//获取缓存中的token信息
func (s *userTokenService) GetUserIdByToken(token string) (userId int64, err error) {
	key := fmt.Sprintf("user:token2uid:%s", token)
	userId, err = gredis.GetRedis().Get(context.TODO(), key).Int64()
	return
}

//根据用户ID获取token
func (s *userTokenService) GetRefreshTokenByUserId(userId int64) string {
	key := fmt.Sprintf("user:uid2token:%d", userId)
	token := gredis.GetRedis().Get(context.TODO(), key).Val()
	return token
}

//保存用户ID->token关联
func (s *userTokenService) SetUserIdToToken(userId int64, token string) error {
	tokenExpireDays := time.Hour * 24 * 40
	key := fmt.Sprintf("user:uid2token:%d", userId)
	err := gredis.GetRedis().Set(context.TODO(), key, token, tokenExpireDays).Err()
	return err
}
