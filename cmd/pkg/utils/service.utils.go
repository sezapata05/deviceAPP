package utilsservices

import (
	modelpackage "dc-nearshore/cmd/pkg/models"
	"time"

	"github.com/rs/zerolog/log"
)

func FormatResponse(status_code int, status bool, message string, results interface{}) modelpackage.ResponseModel {
	response := modelpackage.ResponseModel{
		Status_code: status_code,
		Status:      status,
		Message:     message,
		Results:     results,
	}

	return response
}

func GenerateMapsFromBody(bodyReq modelpackage.DeviceBody) []modelpackage.FirmwareModel {

	firmwares := []modelpackage.FirmwareModel{}

	for _, firmware := range bodyReq.Firmwares {

		releaseDate, err := time.Parse("2006-01-02", firmware.ReleaseDate)
		if err != nil {
			log.Print(err)
			return firmwares
		}

		firmwares = append(firmwares, modelpackage.FirmwareModel{
			ID:           firmware.ID,
			Name:         firmware.FirmwareName,
			ReleaseNotes: firmware.ReleaseNotes,
			ReleaseDate:  releaseDate,
			Url:          firmware.Url,
			CreatedAt:    time.Now(),
		})
	}

	return firmwares
}
