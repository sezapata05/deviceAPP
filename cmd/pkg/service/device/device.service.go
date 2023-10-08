package devicepackage

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"

	modelpackage "dc-nearshore/cmd/pkg/models"
	utilsservices "dc-nearshore/cmd/pkg/utils"
	"time"
)

type deviceService struct {
	DeviceRepository modelpackage.DeviceInternalRepository
}

type DeviceCofig struct {
	DeviceRepository modelpackage.DeviceInternalRepository
}

func NewDeviceService(c *DeviceCofig) modelpackage.DeviceService {
	return &deviceService{
		DeviceRepository: c.DeviceRepository,
	}
}

// CreateDevice implements modelpackage.DeviceService.
func (serviceObj *deviceService) CreateDevice(bodyReq *modelpackage.DeviceBody) modelpackage.ResponseModel {

	firmwares := utilsservices.GenerateMapsFromBody(*bodyReq)

	deviceObj := modelpackage.DeviceModel{
		ID:        uuid.New(),
		Name:      bodyReq.Name,
		Firmware:  firmwares,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}

	objDevice := serviceObj.DeviceRepository.CreateDevice(&deviceObj)
	if _, ok := objDevice.(error); ok {
		return utilsservices.FormatResponse(http.StatusInternalServerError, false, "", objDevice)
	}
	return utilsservices.FormatResponse(http.StatusCreated, true, "", objDevice)
}

// GetAllDevice implements modelpackage.DeviceService.
func (serviceObj *deviceService) GetAllDevice() modelpackage.ResponseModel {

	allDevices, err := serviceObj.DeviceRepository.GetAllDevice()

	if err != nil {
		return utilsservices.FormatResponse(http.StatusNotFound, false, "devices not found", []string{})
	}
	return utilsservices.FormatResponse(http.StatusOK, true, "", allDevices)
}

// GetDeviceById implements modelpackage.DeviceService.
func (serviceObj *deviceService) GetDeviceById(id string) modelpackage.ResponseModel {
	parsedID, errUUid := uuid.Parse(id)

	if errUUid != nil {
		return utilsservices.FormatResponse(http.StatusNotFound, false, "uuid not valid", "uuid not valid")
	}

	deviceID, err := serviceObj.DeviceRepository.GetDeviceById(parsedID)

	if err != nil {
		return utilsservices.FormatResponse(http.StatusNotFound, false, "devices not found", []string{})
	}
	return utilsservices.FormatResponse(http.StatusOK, true, "", deviceID)
}

// UpdateDevice implements modelpackage.DeviceService.
func (serviceObj *deviceService) UpdateDevice(idDevice string, bodyReq *modelpackage.DeviceBody) modelpackage.ResponseModel {

	parsedID, errUUid := uuid.Parse(idDevice)

	if errUUid != nil {
		return utilsservices.FormatResponse(http.StatusNotFound, false, "uuid not valid", "uuid not valid")
	}

	deviceValidated, errValidated := serviceObj.DeviceRepository.GetDeviceById(parsedID)

	if errValidated != nil {
		return utilsservices.FormatResponse(http.StatusNotFound, false, "devices not found", []string{})
	}

	firmwares := utilsservices.GenerateMapsFromBody(*bodyReq)

	deviceObj := modelpackage.DeviceModel{
		ID:        parsedID,
		Name:      bodyReq.Name,
		Firmware:  firmwares,
		CreatedAt: deviceValidated.CreatedAt,
	}

	objDevice := serviceObj.DeviceRepository.UpdateDevice(&deviceObj)

	if _, ok := objDevice.(error); ok {
		return utilsservices.FormatResponse(http.StatusInternalServerError, false, "", objDevice)
	}
	return utilsservices.FormatResponse(http.StatusCreated, true, "", objDevice)

}

// DeleteDevice implements modelpackage.DeviceService.
func (serviceObj *deviceService) DeleteDevice(id string) modelpackage.ResponseModel {
	parsedID, errUUid := uuid.Parse(id)

	if errUUid != nil {
		return utilsservices.FormatResponse(http.StatusNotFound, false, "uuid not valid", "uuid not valid")
	}

	_, errValidated := serviceObj.DeviceRepository.GetDeviceById(parsedID)

	if errValidated != nil {
		return utilsservices.FormatResponse(http.StatusNotFound, false, "devices not found", []string{})
	}

	err := serviceObj.DeviceRepository.DeleteDevice(parsedID)

	if err != nil {
		return utilsservices.FormatResponse(http.StatusBadRequest, false, "", err)
	}
	return utilsservices.FormatResponse(http.StatusOK, true, "", fmt.Sprintf("Device id: %v, deleted", parsedID))
}
