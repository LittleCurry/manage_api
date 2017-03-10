package handle

import (
	"github.com/labstack/echo"
	"strconv"
	"github.com/jinzhu/copier"
	"fmt"
	"net/http"
	"git.iguiyu.com/park/locker-admin-back/model"
	"git.iguiyu.com/park/locker-admin-back/db"
	"git.iguiyu.com/park/locker-admin-back/vm"
)

/**
锁主列表
 */
func GetLockerOwners(c echo.Context) (err error) {
	start, _ := strconv.Atoi(c.Param("start"))
	limit, _ := strconv.Atoi(c.Param("limit"))
	filter := c.QueryParam("usernamefilter")
	var error1, error2 error
	var total int64
	owners := make([]model.OwnerUser, 0)
	engine := db.GetMySQL()
	if filter == "" {
		error1 = engine.Limit(limit, start).Find(&owners)
		total, error2 = engine.Count(new(model.OwnerUser))
	} else {
		error1 = engine.Where("loginname like '%" + filter + "%'").Limit(limit, start).Find(&owners)
		total, error2 = engine.Where("loginname like '%" + filter + "%'").Count(new(model.OwnerUser))
	}

	if error1 == nil && error2 == nil {
		txOwnerUsers := new(vm.TxOwnerUsers)
		copier.Copy(&txOwnerUsers.OwnerUsers, &owners)
		txOwnerUsers.Total = total
		//fmt.Println("lockerUsers:", txOwnerUsers)
		return c.JSON(http.StatusOK, txOwnerUsers)
	} else {
		fmt.Println("锁主列表：空")
		return c.JSON(http.StatusNoContent, nil)
	}
}

/**
车主列表
 */
func GetCarUsers(c echo.Context) (err error) {
	start, _ := strconv.Atoi(c.Param("start"))
	limit, _ := strconv.Atoi(c.Param("limit"))
	filter := c.QueryParam("usernamefilter")

	var error1, error2 error
	var total int64
	carUsers := make([]model.CarUser, 0)
	engine := db.GetMySQL()
	if filter == "" {
		error1 = engine.Limit(limit, start).Find(&carUsers)
		total, error2 = engine.Count(new(model.CarUser))
	} else {
		//如果有查询关键字，此处添加查询方法
	}

	if error1 == nil && error2 == nil {
		txCarUsers := new(vm.TxCarUsers)
		copier.Copy(&txCarUsers.CarUsers, &carUsers)
		txCarUsers.Total = total
		//fmt.Println("carUsers:", txCarUsers)
		return c.JSON(http.StatusOK, txCarUsers)
	} else {
		fmt.Println("车主列表：空")
		return c.JSON(http.StatusNoContent, nil)
	}

}
