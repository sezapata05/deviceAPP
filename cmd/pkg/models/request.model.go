package modelpackage

import "github.com/google/uuid"

type DeviceBody struct {
	Name      string          `json:"device_name"`
	Firmwares []FirmwaresBody `json:"firmwares"`
}

type FirmwaresBody struct {
	ID           uuid.UUID `json:"firmware_id"`
	FirmwareName string    `json:"firmware_name"`
	ReleaseNotes string    `json:"release_notes"`
	ReleaseDate  string    `json:"release_date"`
	Url          string    `json:"url"`
}

type FirmwareBody struct {
	FirmwareName string    `json:"firmware_name"`
	DeviceID     uuid.UUID `json:"device_id"`
	ReleaseNotes string    `json:"release_notes"`
	ReleaseDate  string    `json:"release_date"`
	Url          string    `json:"url"`
}
