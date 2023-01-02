package resultTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/customResponse"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/result/resultBiz"
	"managerstudent/modules/result/resultModel"
	"managerstudent/modules/result/resultStorage"
)

func DeleteResults(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter resultModel.Filter

		if err := c.BindJSON(&filter); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		store := resultStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := resultBiz.NewDeleteResultBiz(store)
		if err := biz.DeleteResult(c.Request.Context(), filter); err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse("success"))
	}
}
