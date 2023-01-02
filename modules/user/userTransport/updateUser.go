package userTransport

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/customResponse"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/component/hasher/Hash_local"
	"managerstudent/modules/user/userBiz"
	"managerstudent/modules/user/userModel"
	"managerstudent/modules/user/userStorage"
)

func UpdateUser(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data userModel.User

		if err := c.ShouldBind(&data); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		filter := bson.D{{"user_name", data.UserName}}

		store := userStorage.NewMongoStore(app.GetNewDataMongoDB())
		md5 := Hash_local.NewHashInfo()
		biz := userBiz.NewUpdateBusiness(store, md5)
		err := biz.UpdateUser(c.Request.Context(), filter, &data)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse(data))
	}
}
