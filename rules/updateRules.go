package rules

import (
	"managerstudent/common/customResponse"
	"managerstudent/common/solveError"
	"managerstudent/component"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateRules(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := Read()
		if err := c.ShouldBind(&data); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}
		Write(data)
		c.JSON(200, customResponse.SimpleSuccessReponse(bson.M{"message": "Update rule success"}))
	}
}
