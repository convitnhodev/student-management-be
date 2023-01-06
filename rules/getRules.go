package rules

import (
	"managerstudent/common/customResponse"
	"managerstudent/component"

	"github.com/gin-gonic/gin"
)

func GetRules(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := Read()
		c.JSON(200, customResponse.SimpleSuccessReponse(data))
	}
}
