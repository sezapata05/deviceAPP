package modelpackage

import (
	"github.com/google/uuid"
)

/*
* Firmware
 */

type FirmwareService interface {
	GetAllFirmwares() ResponseModel
	GetFirmwareById(id string) ResponseModel
	CreateFirmware(firmwareBody *FirmwareBody) ResponseModel
	UpdateFirmware(idFirmware string, firmwareBody *FirmwareBody) ResponseModel
	DeleteFirmware(id string) ResponseModel
}

type FirmwareInternalRepository interface {
	GetAllFirmwares() (*[]FirmwareModel, error)
	GetFirmwareById(id uuid.UUID) (*FirmwareModel, error)
	CreateFirmware(firmwareModel *FirmwareModel) interface{}
	UpdateFirmware(firmwareModel *FirmwareModel) interface{}
	DeleteFirmware(id uuid.UUID) error
}

/*
* Device
 */

type DeviceService interface {
	GetAllDevice() ResponseModel
	GetDeviceById(id string) ResponseModel
	CreateDevice(bodyReq *DeviceBody) ResponseModel
	UpdateDevice(idDevice string, bodyReq *DeviceBody) ResponseModel
	DeleteDevice(id string) ResponseModel
}

type DeviceInternalRepository interface {
	GetAllDevice() (*[]DeviceModel, error)
	GetDeviceById(id uuid.UUID) (*DeviceModel, error)
	CreateDevice(deviceObj *DeviceModel) interface{}
	UpdateDevice(deviceObj *DeviceModel) interface{}
	DeleteDevice(id uuid.UUID) error
}
