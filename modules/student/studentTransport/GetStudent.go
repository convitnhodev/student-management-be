package studentTransport

import (
	"managerstudent/common/customResponse"
	"managerstudent/component"
	"managerstudent/modules/student/studentStorage"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetStudent(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Request.URL.Query().Get("id")

		store := studentStorage.NewMongoStore(app.GetNewDataMongoDB())
		data, err := store.FindStudent(c.Request.Context(), bson.D{{"id", id}})
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse(data))
	}
}
