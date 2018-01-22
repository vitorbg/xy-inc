package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vitorbg/xy-inc/conf"
	"github.com/vitorbg/xy-inc/model"
	"github.com/vitorbg/xy-inc/resources"
	"github.com/vitorbg/xy-inc/util"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {

	err := conf.LoadConf("xy-inc.json")
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(conf.GetUrlDatabase()); os.IsNotExist(err) {
		log.Println(conf.GetUrlDatabase())
		model.InitDatabase()
	}

	if _, err := os.Stat("log"); os.IsNotExist(err) {
		os.Mkdir("log", 0777)
	}

	lf := util.NewLogFile("log" + string(filepath.Separator) + "xy-inc.log")
	log.SetFlags(log.Llongfile + log.Ltime + log.Ldate)
	log.SetOutput(io.MultiWriter(lf, os.Stdout))

	util.RotatesLog(lf)

	e := echo.New()
	e.Logger.SetOutput(io.MultiWriter(lf, os.Stdout))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: lf}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: os.Stdout}))

	e.GET("/poi", resources.FindPoiAll)
	e.POST("/poi", resources.SavePoi)
	e.GET("/poi/proximidade", resources.FindPoiByCoordinateAndDistance)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(conf.GetAppPort())))
}
