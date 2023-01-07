package resultTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/customResponse"
	"managerstudent/component"
	"managerstudent/modules/result/resultBiz"
	"managerstudent/modules/result/resultStorage"
)

func GetAvgResult(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := c.GetQuery("id")
		if ok == false {
			panic("id is not exist")
		}
		store := resultStorage.NewMongoStore(appCtx.GetNewDataMongoDB())
		biz := resultBiz.NewCountAvgMarkBiz(store)
		result, err := biz.CountResult(c.Request.Context(), id)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse(result))
	}
}
