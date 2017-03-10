package handle

import (
	"github.com/labstack/echo"
	"strconv"
	"git.iguiyu.com/park/locker-admin-back/model"
	"git.iguiyu.com/park/locker-admin-back/db"
	"git.iguiyu.com/park/locker-admin-back/vm"
	"github.com/jinzhu/copier"
	"net/http"
	"git.iguiyu.com/park/locker-admin-back/joinmodel"
	"github.com/boombuler/barcode/qr"
	"github.com/boombuler/barcode"
	"image/jpeg"
	"bytes"
)

/**
二维码生成
 */
func GetQRCode(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	engine := db.GetMySQL()
	var locker = model.Locker{Id: id}
	has, err := engine.Get(&locker)
	if (err == nil && has) {
		code, _ := qr.Encode(locker.Qrcode, qr.H, qr.Auto)
		code, _ = barcode.Scale(code, 200, 200)
		var buf bytes.Buffer
		jpeg.Encode(&buf, code, &jpeg.Options{Quality: 75, })
		return c.Blob(http.StatusOK, "image/jpg", buf.Bytes())
	} else {
		return c.Blob(http.StatusNoContent, "image/jpg", nil)
	}
}

/**
出厂地锁列表
 */
func GetLockers(c echo.Context) (err error) {
	start, _ := strconv.Atoi(c.Param("start"))
	limit, _ := strconv.Atoi(c.Param("limit"))
	filter := c.QueryParam("shortidfilter")
	var err1, err2 error
	var total int64

	lockers := make([]model.Locker, 0)
	engine := db.GetMySQL()
	if filter == "" {
		err1 = engine.Limit(limit, start).Find(&lockers)
		total, err2 = engine.Count(new(model.Locker))
	} else {
		err1 = engine.Where("short_id like '%" + filter + "%'").Limit(limit, start).Find(&lockers)
		total, err2 = engine.Where("short_id like '%" + filter + "%'").Count(new(model.Locker))
	}
	if err1 == nil && err2 == nil {
		txLockers := new(vm.TxLockers)
		copier.Copy(&txLockers.Lockers, &lockers)
		txLockers.Total = total
		return c.JSON(http.StatusOK, txLockers)
	} else {
		return c.JSON(http.StatusNoContent, nil)
	}
}

/**
激活地锁列表
 */

func GetActiveLockers(c echo.Context) (err error) {
	start, _ := strconv.Atoi(c.Param("start"))
	limit, _ := strconv.Atoi(c.Param("limit"))
	filter := c.QueryParam("shortidfilter")
	var err1, err2 error
	var total int64

	activeLockers := make([]joinmodel.LockerOwner, 0)
	engine := db.GetMySQL()
	if filter == "" {
		err1 = engine.Join("INNER", "owner_user", "owner_user.id = locker.owner_user_id").Where("active = 1").Limit(limit, start).Find(&activeLockers)
		total, err2 = engine.Where("active = 1").Count(new(model.Locker))

	} else {
		err1 = engine.Join("INNER", "owner_user", "owner_user.id = locker.owner_user_id").Where("active = 1").Where("short_id like '%" + filter + "%'").Limit(limit, start).Find(&activeLockers)
		total, err2 = engine.Where("active = 1").Where("short_id like '%" + filter + "%'").Count(new(model.Locker))
	}
	if err1 == nil && err2 == nil {
		txActiveLockers := new(vm.TxActiveLockers)
		copier.Copy(&txActiveLockers.ActiveLockers, &activeLockers)
		txActiveLockers.Total = total
		//fmt.Println("lockerOwner:", lockers)
		//fmt.Println("activelockers:", txActiveLockers.ActiveLockers)
		return c.JSON(http.StatusOK, txActiveLockers)
	} else {
		//fmt.Println("激活地锁：空")
		return c.JSON(http.StatusNoContent, nil)
	}
}

/**
更新地锁
 */
func UpdateLocker(c echo.Context) (err error) {
	locker := new(vm.RxLocker)
	err = c.Bind(locker)
	engine := db.GetMySQL()
	if err != nil {
		return err
	} else {
		newLocker := new(model.Locker)
		copier.Copy(&newLocker, &locker)
		engine.Id(newLocker.Id).Update(newLocker)
		return c.JSON(http.StatusCreated, locker)
	}

}
