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

func UserUpdatePassword(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data userModel.UpdatePassWord

		if err := c.ShouldBind(&data); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		filter := bson.D{{"user_name", data.UserName}}

		store := userStorage.NewMongoStore(app.GetNewDataMongoDB())
		md5 := Hash_local.NewHashInfo()
		biz := userBiz.NewUpdatepasswordBusiness(store, md5)
		err := biz.UpdatePasswordOfUser(c.Request.Context(), filter, &data)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse("ok"))
	}
}
