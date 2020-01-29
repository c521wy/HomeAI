package auth

import (
	"HomeAI/config"
	"errors"
)

// 认证
func Auth(userToken string) error {
	if userToken == config.AppConfig.AuthToken {
		return nil
	}
	return errors.New("非法的token")
}
