package userTransport

import (
	"managerstudent/common/customResponse"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/user/userStorage"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func AcceptUser(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Request.URL.Query().Get("username")
		store := userStorage.NewMongoStore(app.GetNewDataMongoDB())
		user, err := store.FindUser(c.Request.Context(), bson.M{"username": username})

		if err != nil {
			c.JSON(400, err)
			return
		}

		if user == nil {
			c.JSON(400, solveError.ErrInvalidRequest(nil))
			return
		}

		filter := bson.D{{"username", username}}

		err = store.UpdateUser(c.Request.Context(), filter, bson.M{"acp": true})

		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse(bson.M{"message": "accept user success"}))
	}
}
