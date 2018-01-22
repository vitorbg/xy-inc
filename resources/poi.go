package resources

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/vitorbg/xy-inc/model"
	"log"
	"net/http"
	"strconv"
)

func SavePoi(c echo.Context) error {
	name := c.FormValue("nome")
	coordinate_x := c.FormValue("x")
	coordinate_y := c.FormValue("y")

	x, err := strconv.Atoi(coordinate_x)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Erro no parametro x. Deve ser um numero natural.")
	}

	y, err := strconv.Atoi(coordinate_y)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Erro no parametro y. Deve ser um numero natural.")
	}

	err = model.SavePoi(name, x, y)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Println(fmt.Sprintf("POI[%s,x=%d,y=%d] gravado com sucesso. ", name, x, y))
	return c.JSON(http.StatusCreated, "Poi gravado com sucesso.")
}

func FindPoiAll(c echo.Context) error {

	listPoi, err := model.FindPoiAll()
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Erro ao buscar todos os POIs cadastrados.")
	}

	return c.JSON(http.StatusOK, listPoi)
}

func FindPoiByCoordinateAndDistance(c echo.Context) error {
	coordinate_x := c.QueryParam("x")
	coordinate_y := c.QueryParam("y")
	dmax := c.QueryParam("dmax")

	x, err := strconv.Atoi(coordinate_x)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Erro no parametro x.")
	}

	y, err := strconv.Atoi(coordinate_y)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Erro no parametro y.")
	}

	d, err := strconv.Atoi(dmax)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "Erro no parametro dmax.")
	}

	listPoi, err := model.FindPoiByCoordinateAndDistance(x, y, d)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, listPoi)
}
