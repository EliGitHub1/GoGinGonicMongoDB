package employee_controller

import (
	"net/http"
	"vacation_requests/model"
	"vacation_requests/services"
	"vacation_requests/utils/rest_errors"

	"github.com/gin-gonic/gin"
)

func GetVactionRequest(c *gin.Context) {
	c.String(http.StatusNotImplemented, "GetVactionRequest")
}

func GetRemainingDays(c *gin.Context) {
	c.String(http.StatusNotImplemented, "GetRemainingDays")
}

func CreateVactionRequest(c *gin.Context) {
	var input model.Request
	if err := c.ShouldBindJSON(&input); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(http.StatusInternalServerError, gin.H{"status": restErr.Status(), "des": restErr})
		return
	}
	remDays, remDaysErr := services.GetRemainingDays(input)
	if remDaysErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": remDaysErr.Status(), "des": remDaysErr})
		return
	}

	if remDays > 0 {
		saveErr := services.CreateRequest(input)
		if saveErr != nil {
			c.JSON(saveErr.Status(), gin.H{"Description": saveErr.Description(), "Message": saveErr.Message()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"result": "Created"})
}
