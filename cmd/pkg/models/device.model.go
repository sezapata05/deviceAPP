package modelpackage

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// type DeviceModel struct {
// 	ID        uuid.UUID       `gorm:"column:ID;type:uuid;primary_key" json:"id"`
// 	Name      string          `gorm:"column:Name;not null" json:"name"`
// 	CreatedAt time.Time       `gorm:"column:CreatedAt" json:"created_at"`
// 	UpdatedAt time.Time       `gorm:"column:UpdatedAt" json:"updated_at"`
// 	DeletedAt *gorm.DeletedAt `gorm:"column:DeletedAt;type:date;default:NULL" json:"deleted_at,omitempty"`
// }

type DeviceModel struct {
	ID        uuid.UUID       `gorm:"column:device_id;type:uuid;primary_key;default:gen_random_uuid()" json:"device_id"`
	Name      string          `gorm:"column:name;not null" json:"device_name"`
	Firmware  []FirmwareModel `gorm:"foreignKey:device_id;references:device_id" json:"firmwares,omitempty"`
	CreatedAt time.Time       `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at;type:date;default:NULL" json:"deleted_at,omitempty"`
}

func (*DeviceModel) TableName() string {
	return "public.DEVICE"
}
