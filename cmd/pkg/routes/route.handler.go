package routehandlerpackage

import (
	interfacepackage "dc-nearshore/cmd/pkg/models"

	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"

	middleware "dc-nearshore/cmd/pkg/middleware"
)

var validate = validator.New()
var ve validator.ValidationErrors

type Handler struct {
	DeviceService   interfacepackage.DeviceService
	FirmwareService interfacepackage.FirmwareService
}

type Config struct {
	R               *gin.Engine
	DeviceService   interfacepackage.DeviceService
	FirmwareService interfacepackage.FirmwareService
}

func NewHandler(c *Config) {
	h := &Handler{
		DeviceService:   c.DeviceService,
		FirmwareService: c.FirmwareService,
	}

	// middleware auth for all endpoints
	c.R.Use(middleware.BasicAuthMiddleware())

	ping := c.R.Group("/ping")
	{
		ping.GET("", h.Ping)
	}

	// create a api version
	version := "v1"

	// API name
	dcnearshore := c.R.Group("/dcnearshore/" + version)

	// Device
	devices := dcnearshore.Group("/devices")
	{
		devices.GET("", h.GetAllDeviceController)
		devices.GET(":id", h.GetDeviceController)
		devices.POST("", h.PostDeviceController)
		devices.PUT(":id", h.PutDeviceController)
		devices.DELETE(":id", h.DeleteDeviceController)
	}

	// Firmware
	firmwares := dcnearshore.Group("/firmwares")
	{
		firmwares.GET("", h.GetAllFirmwareController)
		firmwares.GET(":id", h.GetFirmwareController)
		firmwares.POST("", h.PostFirmwareController)
		firmwares.PUT(":id", h.PutFirmwareController)
		firmwares.DELETE(":id", h.DeleteFirmwareController)
	}

}
