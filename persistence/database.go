package persistence

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vitorbg/xy-inc/conf"
	"log"
)

func NewDataBase(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func InitDatabase() error {
	db, err := NewDataBase(conf.GetUrlDatabase())
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

	_, err2 := tx.Exec(`
    CREATE TABLE poi (
      id_poi INTEGER    PRIMARY KEY AUTOINCREMENT NOT NULL,
      nome   STRING     NOT NULL,
      x      INTEGER    NOT NULL,
      y      INTEGER    NOT NULL);`)

	if err2 != nil {
		log.Println(err2)
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
