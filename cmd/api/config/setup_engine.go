package configpackage

import (
	devicerepositorypackage "dc-nearshore/cmd/pkg/repository/device"
	firmwarerepositorypackage "dc-nearshore/cmd/pkg/repository/firmware"
	routehandlerpackage "dc-nearshore/cmd/pkg/routes"

	devicepackage "dc-nearshore/cmd/pkg/service/device"
	firmwarepackage "dc-nearshore/cmd/pkg/service/firmware"

	"github.com/gin-gonic/gin"
)

func SetUp(d *DataSource) (*gin.Engine, error) {

	// fmt.Println("Starting data source")

	/*
	* repository layer
	 */
	device_repository := devicerepositorypackage.NewDeviceRepository(d.DBPostgre)
	firmware_repository := firmwarerepositorypackage.NewFirmwareRepository(d.DBPostgre)

	/*
	* repository layer
	 */
	device_service := devicepackage.NewDeviceService(&devicepackage.DeviceCofig{
		DeviceRepository: device_repository,
	})
	firmware_service := firmwarepackage.NewFirmwareService(&firmwarepackage.FirmwareCofig{
		FirmwareRepository: firmware_repository,
	})

	// Inizialize gin.Engine
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	router := gin.New()

	routehandlerpackage.NewHandler(&routehandlerpackage.Config{
		R:               router,
		DeviceService:   device_service,
		FirmwareService: firmware_service,
	})

	return router, nil
}
