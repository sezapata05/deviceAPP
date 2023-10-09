package configpackage

import (
	postgreconfigpackage "dc-nearshore/cmd/api/config/db"
	modelpackage "dc-nearshore/cmd/pkg/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/rs/zerolog/log"
)

type DataSource struct {
	DBPostgre *gorm.DB
}

func generateMigration(db *gorm.DB) error {

	log.Print("Starting migration . . .")

	err := db.AutoMigrate(&modelpackage.DeviceModel{}, &modelpackage.FirmwareModel{})
	if err != nil {
		return err
	}
	return nil
}

func InitDS() (*DataSource, error) {
	dsn, errURI := postgreconfigpackage.GetDBURI()

	if errURI != nil {
		return nil, errURI
	}

	var db *gorm.DB
	var err error
	maxRetries := 10
	retryInterval := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "public",
				SingularTable: true,
			},
			CreateBatchSize: 100,
		})

		if err == nil {
			break
		}

		log.Printf("Error connecting to the database (remaining attempts): %d): %v\n", maxRetries-i-1, err)
		time.Sleep(retryInterval)
	}

	if err != nil {
		return nil, err
	}

	log.Print("Postgre data base connected")

	errMigration := generateMigration(db)

	if errMigration != nil {
		return nil, err
	}

	return &DataSource{
		DBPostgre: db,
	}, nil
}
