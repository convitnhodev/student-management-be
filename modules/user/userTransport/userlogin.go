package userTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/customResponse"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/component/hasher/Hash_local"
	"managerstudent/component/tokenProvider/jwt"
	"managerstudent/modules/user/userBiz"
	"managerstudent/modules/user/userModel"
	"managerstudent/modules/user/userStorage"
	"net/http"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData userModel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			solveError.ErrInvalidRequest(err)
		}

		db := appCtx.GetNewDataMongoDB()

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecret())

		store := userStorage.NewMongoStore(db)
		hasherInfo := Hash_local.NewHashInfo()
		biz := userBiz.NewLoginBusiness(store, tokenProvider, hasherInfo, 60*60*24*30)

		account, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			c.JSON(400, err)
			return
		}

		c.JSON(http.StatusOK, customResponse.SimpleSuccessReponse(account))

	}
}
