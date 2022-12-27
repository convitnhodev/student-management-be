package userTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/customResponse"
	"managerstudent/common/paging"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/user/userBiz"
	"managerstudent/modules/user/userModel"
	"managerstudent/modules/user/userStorage"
)

func ListUsers(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var page paging.Paging
		if err := c.ShouldBind(&page); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		var filter userModel.Filter
		if err := c.BindJSON(&filter); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		store := userStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := userBiz.NewListUsersBiz(store)
		data, err := biz.ListUsers(c.Request.Context(), filter, &page)
		if err != nil {
			c.JSON(400, err)
			return
		}
		page.Fullfill()
		c.JSON(200, customResponse.NewSuccessReponse(data, page, nil))
	}
}
