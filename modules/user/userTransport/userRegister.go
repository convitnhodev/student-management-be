package userTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/customResponse"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/component/hasher/Hash_local"
	"managerstudent/modules/user/userBiz"
	"managerstudent/modules/user/userModel"
	"managerstudent/modules/user/userStorage"
)

func UserRegister(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data userModel.User
		if err := c.ShouldBind(&data); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}
		//
		store := userStorage.NewMongoStore(app.GetNewDataMongoDB())
		md5 := Hash_local.NewHashInfo()
		biz := userBiz.NewCreateUserBiz(store, md5, app.GetPubsub())
		if err := biz.CreateNewUser(c.Request.Context(), &data); err != nil {
			c.JSON(400, err)
			return
		}
		dataReturn := map[string]interface{}{"user_name": data.UserName, "role": data.Role}
		c.JSON(200, customResponse.SimpleSuccessReponse(dataReturn))
	}
}
