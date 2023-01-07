package userTransport

import (
	"github.com/gin-gonic/gin"
	_const "managerstudent/common/const"
	"managerstudent/common/customResponse"
	"managerstudent/component"
	"net/http"
)

func GetProfile(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(_const.CurrentUser).(_const.Requester)
		c.JSON(http.StatusOK, customResponse.SimpleSuccessReponse(data.GetUserName()))
	}
}
