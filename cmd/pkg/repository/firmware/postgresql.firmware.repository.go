package firmwarerepositorypackage

import (
	modelpackage "dc-nearshore/cmd/pkg/models"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type postgresqlRepository struct {
	DB *gorm.DB
}

// CreateFirmware implements modelpackage.FirmwareInternalRepository.
func (dbrepositoryPKG *postgresqlRepository) CreateFirmware(firmwareModel *modelpackage.FirmwareModel) interface{} {

	tx := dbrepositoryPKG.DB.Begin()

	result := tx.Create(&firmwareModel)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()

	return &firmwareModel
}

// DeleteFirmware implements modelpackage.FirmwareInternalRepository.
func (dbrepositoryPKG *postgresqlRepository) DeleteFirmware(id uuid.UUID) error {

	var firmware modelpackage.FirmwareModel

	result := dbrepositoryPKG.DB.Where("firmware_id = ?", id).Delete(&firmware)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetAllFirmwares implements modelpackage.FirmwareInternalRepository.
func (dbrepositoryPKG *postgresqlRepository) GetAllFirmwares() (*[]modelpackage.FirmwareModel, error) {

	var firmwares []modelpackage.FirmwareModel

	result := dbrepositoryPKG.DB.Find(&firmwares)

	if result.RowsAffected == 0 || result.Error != nil {
		return nil, fmt.Errorf("could not get firmwares, error: %v", result.Error)
	}

	return &firmwares, nil
}

// GetFirmwareById implements modelpackage.FirmwareInternalRepository.
func (dbrepositoryPKG *postgresqlRepository) GetFirmwareById(id uuid.UUID) (*modelpackage.FirmwareModel, error) {

	var firmware modelpackage.FirmwareModel

	result := dbrepositoryPKG.DB.Where("firmware_id = ?", id).First(&firmware)

	if result.Error != nil {
		return nil, fmt.Errorf("error finding firmware: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("firmware not found: %v", id)
	}

	return &firmware, nil
}

// UpdateFirmware implements modelpackage.FirmwareInternalRepository.
func (dbrepositoryPKG *postgresqlRepository) UpdateFirmware(firmwareModel *modelpackage.FirmwareModel) interface{} {

	tx := dbrepositoryPKG.DB.Begin()

	result := tx.Save(&firmwareModel)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()

	return &firmwareModel
}

func NewFirmwareRepository(db *gorm.DB) modelpackage.FirmwareInternalRepository {
	return &postgresqlRepository{
		DB: db,
	}
}
