package classTransport

import (
	"managerstudent/common/customResponse"
	"managerstudent/component"
	"managerstudent/modules/class/classStorage"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetClass(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		classId := c.Request.URL.Query().Get("class_id")

		store := classStorage.NewMongoStore(app.GetNewDataMongoDB())
		data, err := store.FindClass(c.Request.Context(), bson.M{"class_id": classId})
		if err != nil {
			c.JSON(400, err)
			return
		}

		c.JSON(200, customResponse.SimpleSuccessReponse(data))
	}
}
