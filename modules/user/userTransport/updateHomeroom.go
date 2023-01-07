package userTransport

import (
	"managerstudent/common/customResponse"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/user/userModel"
	"managerstudent/modules/user/userStorage"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateHomeroom(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data userModel.UpdateHomeroom

		if err := c.ShouldBind(&data); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		filter := bson.D{{"username", data.UserName}}

		store := userStorage.NewMongoStore(app.GetNewDataMongoDB())

		_, err := store.FindUser(c.Request.Context(), bson.M{"username": data.UserName})
		if err != nil {
			c.JSON(400, err)
			return
		}

		err = store.UpdateUser(c.Request.Context(), filter, &data)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse(bson.M{"message": "Update homeroom success"}))
	}
}
