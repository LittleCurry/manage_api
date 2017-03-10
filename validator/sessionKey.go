package validator

import (
	"github.com/labstack/echo"
	"git.iguiyu.com/park/locker-admin-back/db"
)

func GetValidator(key string, _ echo.Context) bool {
	_, err := db.GetRedis().Get("SessionKey-" + key).Result()
	if err == nil {
		return true
	} else {
		return false
	}
}
