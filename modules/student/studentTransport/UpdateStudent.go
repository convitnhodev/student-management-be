package studentTransport

import (
	"managerstudent/common/customResponse"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/student/studentModel"
	"managerstudent/modules/student/studentStorage"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateStudent(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data studentModel.Student

		if err := c.ShouldBind(&data); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		filter := bson.D{{"id", data.Id}}

		store := studentStorage.NewMongoStore(app.GetNewDataMongoDB())
		err := store.UpdateStudent(c.Request.Context(), filter, data)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse("success"))
	}
}
