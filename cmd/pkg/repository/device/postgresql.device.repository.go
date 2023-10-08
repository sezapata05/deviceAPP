package devicerepositorypackage

import (
	modelpackage "dc-nearshore/cmd/pkg/models"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type postgresqlRepository struct {
	DB *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) modelpackage.DeviceInternalRepository {
	return &postgresqlRepository{
		DB: db,
	}
}

// CreateDevice implements modelpackage.DeviceInternalRepository.
func (dbrepositoryObj *postgresqlRepository) CreateDevice(deviceObj *modelpackage.DeviceModel) interface{} {

	tx := dbrepositoryObj.DB.Begin()

	result := tx.Create(&deviceObj)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()

	return &deviceObj
}

// GetAllDevice implements modelpackage.DeviceInternalRepository.
func (dbrepositoryObj *postgresqlRepository) GetAllDevice() (*[]modelpackage.DeviceModel, error) {

	var devices []modelpackage.DeviceModel

	result := dbrepositoryObj.DB.
		Preload("Firmware").
		Find(&devices)

	if result.RowsAffected == 0 || result.Error != nil {
		return nil, fmt.Errorf("could not get devices, error: %v", result.Error)
	}

	return &devices, nil
}

// GetDeviceById implements modelpackage.DeviceInternalRepository.
func (dbrepositoryObj *postgresqlRepository) GetDeviceById(id uuid.UUID) (*modelpackage.DeviceModel, error) {

	var device modelpackage.DeviceModel

	result := dbrepositoryObj.DB.Where("device_id = ?", id).
		Preload("Firmware").
		First(&device)

	if result.Error != nil {
		return nil, fmt.Errorf("error finding device: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("device not found: %v", id)
	}

	return &device, nil
}

// UpdateDevice implements modelpackage.DeviceInternalRepository.
func (dbrepositoryObj *postgresqlRepository) UpdateDevice(deviceObj *modelpackage.DeviceModel) interface{} {

	tx := dbrepositoryObj.DB.Begin()

	// result := tx.Save(&deviceObj)
	result := tx.Model((&deviceObj)).Updates(&deviceObj)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()

	return &deviceObj
}

// DeleteDevice implements modelpackage.DeviceInternalRepository.
func (dbrepositoryObj *postgresqlRepository) DeleteDevice(id uuid.UUID) error {

	var device modelpackage.DeviceModel

	result := dbrepositoryObj.DB.Where("device_id = ?", id).Delete(&device)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
