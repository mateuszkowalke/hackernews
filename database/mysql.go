package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var Db *sql.DB

func InitDb() {
	db, err := sql.Open("mysql", "mateuszkowalke'Kowalkem2@'@tcp(localhost)/hackernews")
	if err != nil {
		log.Panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	Db = db
}

func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Panic(err)
	}
	driver, _ := mysql.WithInstance(Db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://database/migrations/mysql",
		"mysql",
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
