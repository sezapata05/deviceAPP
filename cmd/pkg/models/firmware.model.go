package modelpackage

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// type FirmwareModel struct {
// 	ID           uuid.UUID       `gorm:"column:ID;type:uuid;primary_key" json:"id"`
// 	Name         string          `gorm:"column:Name;not null" json:"name"`
// 	DeviceID     string          `gorm:"column:DeviceID;not null" json:"device_id"`
// 	ReleaseNotes string          `gorm:"ReleaseNotes:url" json:"release_notes"`
// 	ReleaseDate  time.Time       `gorm:"column:ReleaseDate;type:date" json:"release_date"`
// 	Url          string          `gorm:"column:url" json:"url"`
// 	CreatedAt    time.Time       `gorm:"column:CreatedAt" json:"created_at"`
// 	UpdatedAt    time.Time       `gorm:"column:UpdatedAt" json:"updated_at"`
// 	DeletedAt    *gorm.DeletedAt `gorm:"column:DeletedAt;type:date;default:NULL" json:"deleted_at,omitempty"`
// }

type FirmwareModel struct {
	ID           uuid.UUID       `gorm:"column:firmware_id;type:uuid;primary_key;default:gen_random_uuid()" json:"firmware_id"`
	Name         string          `gorm:"column:name;not null" json:"firmware_name"`
	DeviceID     uuid.UUID       `gorm:"column:device_id;not null" json:"device_id"`
	ReleaseNotes string          `gorm:"column:release_notes" json:"release_notes"`
	ReleaseDate  time.Time       `gorm:"column:release_date;type:date" json:"release_date"`
	Url          string          `gorm:"column:url" json:"url"`
	CreatedAt    time.Time       `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    *time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    *gorm.DeletedAt `gorm:"column:deleted_at;type:date;default:NULL" json:"deleted_at,omitempty"`
}

func (*FirmwareModel) TableName() string {
	return "FIRMWARE"
}
