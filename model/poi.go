package model

import (
	"errors"
	"fmt"
	"github.com/vitorbg/xy-inc/domain"
	"github.com/vitorbg/xy-inc/persistence/dao"
	"log"
	"math"
)

func SavePoi(name string, x int, y int) error {
	if x < 0 {
		return errors.New(fmt.Sprintf("Erro ao gravar POI[%s,x=%d,y=%d]. Parametro x deve ser positivo. ", name, x, y))
	}
	if y < 0 {
		return errors.New(fmt.Sprintf("Erro ao gravar POI[%s,x=%d,y=%d]. Parametro y deve ser positivo. ", name, x, y))
	}
	return dao.SavePoi(name, x, y)
}

func FindPoiAll() ([]domain.Poi, error) {
	listaPoi, err := dao.FindPoiAll()
	return listaPoi, err
}

func FindPoiByCoordinateAndDistance(coordinate_x int, coordinate_y int, distanceMax int) ([]domain.Poi, error) {
	listaPoi, err := dao.FindPoiAll()
	if err != nil {
		log.Println(err)
		return listaPoi, err
	}

	listaPoiAux := make([]domain.Poi, 0)

	for _, poi := range listaPoi {
		distance := calculateDistance(poi.X, poi.Y, coordinate_x, coordinate_y)
		if distance <= float64(distanceMax) {
			listaPoiAux = append(listaPoiAux, poi)
		}
	}

	return listaPoiAux, nil
}

func calculateDistance(x1 int, y1 int, x2 int, y2 int) float64 {
	first := math.Pow(float64(x2-x1), 2)
	second := math.Pow(float64(y2-y1), 2)
	distance := math.Sqrt(first + second)

	return distance
}
