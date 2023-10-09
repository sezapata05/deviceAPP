package routehandlerpackage

import (
	modelpackage "dc-nearshore/cmd/pkg/models"
	validationspackage "dc-nearshore/cmd/pkg/models/validations"
	utilsservices "dc-nearshore/cmd/pkg/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostDeviceController(c *gin.Context) {

	var body modelpackage.DeviceBody

	if err := c.ShouldBindJSON(&body); err != nil {
		response := utilsservices.FormatResponse(http.StatusBadRequest, false, "json has not been sent or with problems", []string{})
		c.JSON(response.Status_code, response)
		return
	}

	if err := validate.Struct(body); err != nil {
		out := validationspackage.ValidationError(err, ve)
		response := utilsservices.FormatResponse(http.StatusBadRequest, false,
			fmt.Sprintf("some error in struct - body: %v", out), []string{})
		c.JSON(response.Status_code, response)
		return
	}

	respCreate := h.DeviceService.CreateDevice(&body)
	c.JSON(respCreate.Status_code, respCreate)

}

func (h *Handler) GetAllDeviceController(c *gin.Context) {

	devicesResults := h.DeviceService.GetAllDevice()
	c.JSON(devicesResults.Status_code, devicesResults)
}

func (h *Handler) GetDeviceController(c *gin.Context) {

	idDevice := c.Param("id")

	result := h.DeviceService.GetDeviceById(idDevice)

	c.JSON(result.Status_code, result)

}

func (h *Handler) PutDeviceController(c *gin.Context) {

	idDevice := c.Param("id")

	var body modelpackage.DeviceBody

	if err := c.ShouldBindJSON(&body); err != nil {
		response := utilsservices.FormatResponse(http.StatusBadRequest, false, "json has not been sent or with problems", []string{})
		c.JSON(response.Status_code, response)
		return
	}

	if err := validate.Struct(body); err != nil {
		out := validationspackage.ValidationError(err, ve)
		response := utilsservices.FormatResponse(http.StatusBadRequest, false,
			fmt.Sprintf("some error in struct - body: %v", out), []string{})
		c.JSON(response.Status_code, response)
		return
	}

	result := h.DeviceService.UpdateDevice(idDevice, &body)

	c.JSON(result.Status_code, result)

}

func (h *Handler) DeleteDeviceController(c *gin.Context) {

	idDevice := c.Param("id")

	result := h.DeviceService.DeleteDevice(idDevice)

	c.JSON(result.Status_code, result)

}
