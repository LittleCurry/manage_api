package handle

import (
	"github.com/labstack/echo"
	"git.iguiyu.com/park/misc/globals"
	"net/http"
	"git.iguiyu.com/park/locker-admin-back/vm"
	"git.iguiyu.com/park/locker-admin-back/model"
	"git.iguiyu.com/park/locker-admin-back/db"
)

/**
	登录相关的请求处理
 */
func Login(c echo.Context) (err error) {
	user := new(vm.RxUser)
	err = c.Bind(user)
	if err != nil {
		return err
	} else {
		var adminUser = model.AdminUser{Loginname: user.LoginName, Password: user.Password}
		has, err := db.GetMySQL().Get(&adminUser)
		if has && err == nil {
			txUser := new(vm.TxUser)
			txUser.LoginName = user.LoginName
			txUser.SessionKey = globals.GenerateSession()
			db.GetRedis().Set("SessionKey-"+txUser.SessionKey, txUser.LoginName, 0)
			return c.JSON(http.StatusOK, txUser)
		} else {
			return c.JSON(http.StatusNoContent, nil)
		}

	}
}
