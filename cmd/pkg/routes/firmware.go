package routehandlerpackage

import (
	modelpackage "dc-nearshore/cmd/pkg/models"
	utilsservices "dc-nearshore/cmd/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostFirmwareController(c *gin.Context) {

	var body modelpackage.FirmwareBody

	if err := c.ShouldBindJSON(&body); err != nil {
		response := utilsservices.FormatResponse(http.StatusBadRequest, false, "json has not been sent or with problems", []string{})
		c.JSON(response.Status_code, response)
		return
	}

	respCreate := h.FirmwareService.CreateFirmware(&body)
	c.JSON(respCreate.Status_code, respCreate)
}

func (h *Handler) GetFirmwareController(c *gin.Context) {

	idFirmware := c.Param("id")

	result := h.FirmwareService.GetFirmwareById(idFirmware)

	c.JSON(result.Status_code, result)
}

func (h *Handler) GetAllFirmwareController(c *gin.Context) {

	firmwareResults := h.FirmwareService.GetAllFirmwares()
	c.JSON(firmwareResults.Status_code, firmwareResults)
}

func (h *Handler) PutFirmwareController(c *gin.Context) {

	idFirmware := c.Param("id")

	var body modelpackage.FirmwareBody

	if err := c.ShouldBindJSON(&body); err != nil {
		response := utilsservices.FormatResponse(http.StatusBadRequest, false, "json has not been sent or with problems", []string{})
		c.JSON(response.Status_code, response)
		return
	}

	result := h.FirmwareService.UpdateFirmware(idFirmware, &body)

	c.JSON(result.Status_code, result)
}

func (h *Handler) DeleteFirmwareController(c *gin.Context) {

	idFirmware := c.Param("id")

	result := h.FirmwareService.DeleteFirmware(idFirmware)

	c.JSON(result.Status_code, result)
}
