package model

import (
	"errors"
	"fmt"
	"github.com/vitorbg/xy-inc/domain"
	"github.com/vitorbg/xy-inc/persistence/dao"
	"math"
)

func SavePoi(name string, x int, y int) error {
	if x < 0 {
		return errors.New(fmt.Sprintf("Erro ao gravar POI[%s,x=%d,y=%d]. Parametro x deve ser positivo. ", name, x, y))
	}
	if y < 0 {
		return errors.New(fmt.Sprintf("Erro ao gravar POI[%s,x=%d,y=%d]. Parametro y deve ser positivo. ", name, x, y))
	}
	err := dao.SavePoi(name, x, y)
	if err != nil {
		return errors.New(fmt.Sprintf("Erro ao gravar POI[%s,x=%d,y=%d]. %s ", name, x, y, err.Error()))
	}
	return nil
}

func FindPoiAll() ([]domain.Poi, error) {
	listPoi, err := dao.FindPoiAll()
	return listPoi, err
}

func FindPoiByCoordinateAndDistance(coordinate_x int, coordinate_y int, distanceMax int) ([]domain.Poi, error) {
	if coordinate_x < 0 {
		return nil, errors.New(fmt.Sprintf("Erro ao buscar POIs por proximidade. Parametro x deve ser positivo. "))
	}
	if coordinate_y < 0 {
		return nil, errors.New(fmt.Sprintf("Erro ao buscar POIs por proximidade. Parametro y deve ser positivo. "))
	}
	if distanceMax < 0 {
		return nil, errors.New(fmt.Sprintf("Erro ao buscar POIs por proximidade. Parametro distancia maxima deve ser positivo. "))
	}

	listPoi, err := dao.FindPoiAll()
	if err != nil {
		return listPoi, errors.New(fmt.Sprintf("Erro ao buscar POIs por proximidade. %s ", err.Error()))
	}

	listPoiAux := make([]domain.Poi, 0)

	for _, poi := range listPoi {
		distance := calculateDistance(poi.X, poi.Y, coordinate_x, coordinate_y)
		if distance <= float64(distanceMax) {
			listPoiAux = append(listPoiAux, poi)
		}
	}

	return listPoiAux, nil
}

func calculateDistance(x1 int, y1 int, x2 int, y2 int) float64 {
	first := math.Pow(float64(x2-x1), 2)
	second := math.Pow(float64(y2-y1), 2)
	distance := math.Sqrt(first + second)

	return distance
}
