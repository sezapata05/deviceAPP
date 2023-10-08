package firmwarepackage

import (
	modelpackage "dc-nearshore/cmd/pkg/models"
	utilsservices "dc-nearshore/cmd/pkg/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type firmwareService struct {
	FirmwareRepository modelpackage.FirmwareInternalRepository
}

type FirmwareCofig struct {
	FirmwareRepository modelpackage.FirmwareInternalRepository
}

func NewFirmwareService(c *FirmwareCofig) modelpackage.FirmwareService {
	return &firmwareService{
		FirmwareRepository: c.FirmwareRepository,
	}
}

// CreateFirmware implements modelpackage.FirmwareService.
func (firmwarePKG *firmwareService) CreateFirmware(firmwareBody *modelpackage.FirmwareBody) modelpackage.ResponseModel {

	releaseDate, _ := time.Parse("2006-01-02", firmwareBody.ReleaseDate)

	firmwareObj := modelpackage.FirmwareModel{
		ID:           uuid.New(),
		Name:         firmwareBody.FirmwareName,
		DeviceID:     firmwareBody.DeviceID,
		ReleaseNotes: firmwareBody.ReleaseNotes,
		ReleaseDate:  releaseDate,
		Url:          firmwareBody.Url,
		CreatedAt:    time.Now(),
		UpdatedAt:    nil,
	}

	objFirmware := firmwarePKG.FirmwareRepository.CreateFirmware(&firmwareObj)

	if _, ok := objFirmware.(error); ok {
		return utilsservices.FormatResponse(http.StatusBadRequest, false, "", objFirmware)
	}
	return utilsservices.FormatResponse(http.StatusCreated, true, "", objFirmware)
}

// DeleteFirmware implements modelpackage.FirmwareService.
func (firmwarePKG *firmwareService) DeleteFirmware(id string) modelpackage.ResponseModel {
	parsedID, errUUid := uuid.Parse(id)

	if errUUid != nil {
		return utilsservices.FormatResponse(http.StatusBadRequest, false, "uuid not valid", "uuid not valid")
	}

	_, errValidate := firmwarePKG.FirmwareRepository.GetFirmwareById(parsedID)

	if errValidate != nil {
		return utilsservices.FormatResponse(http.StatusNotFound, false, "firmware not found", []string{})
	}

	err := firmwarePKG.FirmwareRepository.DeleteFirmware(parsedID)

	if err != nil {
		return utilsservices.FormatResponse(http.StatusBadRequest, false, "", err)
	}
	return utilsservices.FormatResponse(http.StatusOK, true, "", fmt.Sprintf("Firmware id: %v, deleted", parsedID))
}

// GetAllFirmwares implements modelpackage.FirmwareService.
func (firmwarePKG *firmwareService) GetAllFirmwares() modelpackage.ResponseModel {

	allFirmwares, err := firmwarePKG.FirmwareRepository.GetAllFirmwares()

	if err != nil {
		return utilsservices.FormatResponse(http.StatusNotFound, true, "devices not found", []string{})
	}
	return utilsservices.FormatResponse(http.StatusOK, true, "", allFirmwares)
}

// UpdateFirmware implements modelpackage.FirmwareService.
func (firmwarePKG *firmwareService) UpdateFirmware(idFirmware string, firmwareBody *modelpackage.FirmwareBody) modelpackage.ResponseModel {

	parsedID, errUUid := uuid.Parse(idFirmware)

	if errUUid != nil {
		return utilsservices.FormatResponse(http.StatusNotFound, false, "uuid not valid", "uuid not valid")
	}

	firmwareValidated, errValidate := firmwarePKG.FirmwareRepository.GetFirmwareById(parsedID)

	if errValidate != nil {
		return utilsservices.FormatResponse(http.StatusNotFound, false, "firmware not found", []string{})
	}

	releaseDate, _ := time.Parse("2006-01-02", firmwareBody.ReleaseDate)

	firmwareObj := modelpackage.FirmwareModel{
		ID:           parsedID,
		Name:         firmwareBody.FirmwareName,
		DeviceID:     firmwareBody.DeviceID,
		ReleaseNotes: firmwareBody.ReleaseNotes,
		ReleaseDate:  releaseDate,
		Url:          firmwareBody.Url,
		CreatedAt:    firmwareValidated.CreatedAt,
	}

	objFirmware := firmwarePKG.FirmwareRepository.UpdateFirmware(&firmwareObj)

	if _, ok := objFirmware.(error); ok {
		return utilsservices.FormatResponse(http.StatusBadRequest, false, "", objFirmware)
	}
	return utilsservices.FormatResponse(http.StatusOK, true, "", objFirmware)
}

// GetFirmwareById implements modelpackage.FirmwareService.
func (firmwarePKG *firmwareService) GetFirmwareById(id string) modelpackage.ResponseModel {

	parsedID, errUUid := uuid.Parse(id)

	if errUUid != nil {
		return utilsservices.FormatResponse(http.StatusNotFound, false, "uuid not valid", "uuid not valid")
	}

	firmwareID, err := firmwarePKG.FirmwareRepository.GetFirmwareById(parsedID)

	if err != nil {
		return utilsservices.FormatResponse(http.StatusNotFound, false, "firmware not found", []string{})
	}

	return utilsservices.FormatResponse(http.StatusOK, true, "", firmwareID)
}
