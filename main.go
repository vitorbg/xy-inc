package main

import (
	"fmt"
	"github.com/vitorbg/xy-inc/conf"
	"github.com/vitorbg/xy-inc/model"
	"github.com/vitorbg/xy-inc/util"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//os.Remove(conf.GetUrlDatabase())

	err := conf.LoadConf("xy-inc.json")
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(conf.GetUrlDatabase()); os.IsNotExist(err) { //Verifica se a pasta log existe
		log.Println(conf.GetUrlDatabase())
		model.InitDatabase()
	}

	if _, err := os.Stat("log"); os.IsNotExist(err) { //Verifica se a pasta log existe
		os.Mkdir("log", 0777)
	}

	lf := util.NewLogFile("log" + string(filepath.Separator) + "xy-inc.log")
	log.SetOutput(io.MultiWriter(lf, os.Stdout))
	//log.SetFlags(log.Llongfile + log.Ltime + log.Ldate)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	util.RotatesLog(lf)

	log.Println("é nois que voa")
	listaPoi, err := model.FindPoiAll()
	if err != err {
		log.Println(err)
	}
	for _, poi := range listaPoi {
		log.Println(fmt.Sprintf("POI %s - X=%d,Y=%d", poi.Nome, poi.X, poi.Y))
	}

	//	model.FindPoiByCoordinateAndDistance(10, 20, 10)

	//	model.SavePoi("Camelodromo", 50, 58)
}
