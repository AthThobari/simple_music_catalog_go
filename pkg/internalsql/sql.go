package internalsql

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatalf("error to connecting to database %+v\n", err)
		return nil, err
	}
	return db, nil
}