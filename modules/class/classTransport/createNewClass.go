package classTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/class/classBiz"
	"managerstudent/modules/class/classModel"
	"managerstudent/modules/class/classStorage"
)

func CreateNewClass(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data classModel.Class

		if err := c.ShouldBind(&data); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}
		
		store := classStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := classBiz.NewCreateClassBiz(store)
		if err := biz.CreateNewClass(c.Request.Context(), &data); err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, data.Id)
	}

}
