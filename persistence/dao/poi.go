package dao

import (
	"github.com/vitorbg/xy-inc/conf"
	"github.com/vitorbg/xy-inc/domain"
	"github.com/vitorbg/xy-inc/persistence"
	"log"
)

func SavePoi(nome string, x int, y int) error {

	db, err := persistence.NewDataBase(conf.GetUrlDatabase())
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err2 := tx.Exec(`INSERT INTO poi (nome, x, y) VALUES ($1, $2, $3);`, nome, x, y)

	if err2 != nil {
		log.Println(err2)
		tx.Rollback()
		return err2
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func FindPoiAll() ([]domain.Poi, error) {
	var listaPoi []domain.Poi = make([]domain.Poi, 0)

	db, err := persistence.NewDataBase(conf.GetUrlDatabase())
	if err != nil {
		log.Println(err)
		return listaPoi, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT *
		FROM poi;`)
	if err != nil {
		log.Println(err)
		return listaPoi, err
	}

	defer rows.Close()
	for rows.Next() {
		var poi domain.Poi
		err := rows.Scan(&poi.IdPoi,
			&poi.Nome,
			&poi.X,
			&poi.Y)
		if err != nil {
			log.Println(err)
			return listaPoi, err
		}
		listaPoi = append(listaPoi, poi)
	}

	return listaPoi, nil
}
