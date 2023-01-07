package userTransport

import (
	"managerstudent/common/customResponse"
	"managerstudent/component"
	"managerstudent/modules/user/userStorage"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetByUsername(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Request.URL.Query().Get("username")
		store := userStorage.NewMongoStore(app.GetNewDataMongoDB())

		user, err := store.FindUser(c.Request.Context(), bson.M{"username": username})

		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse(user))
	}
}
