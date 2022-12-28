package resultTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/result/resultBiz"
	"managerstudent/modules/result/resultModel"
	"managerstudent/modules/result/resultStorage"
)

func AddResult(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []resultModel.Result
		if err := c.ShouldBind(&result); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		store := resultStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := resultBiz.NewAddResultBiz(store)
		if err := biz.AddResult(c.Request.Context(), result); err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, "ok")
	}
}
