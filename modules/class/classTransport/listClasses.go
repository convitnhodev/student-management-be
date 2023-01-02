package classTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/customResponse"
	"managerstudent/common/paging"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/class/classBiz"
	"managerstudent/modules/class/classModel"
	"managerstudent/modules/class/classStorage"
)

func ListClasses(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var page paging.Paging
		if err := c.ShouldBind(&page); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		var filter classModel.Filter
		if err := c.BindJSON(&filter); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		store := classStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := classBiz.NewListClassesBiz(store)
		data, err := biz.ListClasses(c.Request.Context(), filter, &page)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse(data))
	}
}
