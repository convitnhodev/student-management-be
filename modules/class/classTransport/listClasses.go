package classTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/class/classBiz"
	"managerstudent/modules/class/classStorage"
)

func ListClasses(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var grade int
		if err := c.ShouldBind(&grade); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		store := classStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := classBiz.NewDeleteClassBiz(store)
		err := biz.DeleteClass(c.Request.Context(), &id)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, "success")
	}
}
