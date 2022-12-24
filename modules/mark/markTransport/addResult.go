package markTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/mark/markBiz"
	"managerstudent/modules/mark/markModel"
	"managerstudent/modules/mark/markStorage"
)

func AddResult(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []markModel.Result
		if err := c.ShouldBind(&result); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		store := markStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := markBiz.NewAddResultBiz(store)
		if err := biz.AddResult(c.Request.Context(), result); err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, "ok")
	}
}
