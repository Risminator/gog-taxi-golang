package datastore

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		// data source name, refer https://github.com/jackc/pgx
		DSN: "host=localhost user=postgres password=qwerty123 dbname=gog port=5432 sslmode=disable TimeZone=Europe/Moscow",
		// disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
		PreferSimpleProtocol: true,
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
