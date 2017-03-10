package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"git.iguiyu.com/park/locker-admin-back/handle"
	"github.com/labstack/echo"
	"git.iguiyu.com/park/locker-admin-back/db"
	"github.com/labstack/echo/middleware"
	"git.iguiyu.com/park/locker-admin-back/validator"
)

func main() {
	db.InitMySQL()
	db.InitRedis()
	fmt.Println("Hello World")
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAuthorization},
	}))
	//登录路由
	e.POST("/login", handle.Login)
	//二维码生产路由
	e.GET("/getqrcode/:id", handle.GetQRCode, middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{KeyLookup: "query:sessionkey", Validator: validator.GetValidator, }))

	admin := e.Group("/admin", middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{Validator: validator.GetValidator}))
	{
		//出厂地锁列表路由
		admin.GET("/locker/list/:start/:limit", handle.GetLockers)
		//已激活地锁列表路由
		admin.GET("/locker/listactive/:start/:limit", handle.GetActiveLockers)
		//更新地锁
		admin.POST("/locker/update/", handle.UpdateLocker)
		//锁主列表
		admin.GET("/owner/list/:start/:limit", handle.GetLockerOwners)
		//车主列表
		admin.GET("/driver/list/:start/:limit", handle.GetCarUsers)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
