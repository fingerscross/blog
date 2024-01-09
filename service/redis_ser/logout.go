package redis_ser

import (
	"gvb_server/global"
	"gvb_server/utils"
	"time"
)

// 针对注销的操作
const prefix = "logout_"

// Logout 针对注销的操作
func Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(prefix+token, "", diff).Err()
	return err
}

// CheckLogout 判断是否在redis中 在的话就是注销过
// logout_开头的都找到然后遍历
func CheckLogout(token string) bool {
	keys := global.Redis.Keys(prefix + "*").Val()
	if utils.InList(prefix+token, keys) {
		return true
	}
	return false
}
