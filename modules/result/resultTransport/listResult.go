package resultTransport

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/customResponse"
	"managerstudent/common/paging"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/result/resultBiz"
	"managerstudent/modules/result/resultStorage"
)

func ListResult(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var page paging.Paging
		if err := c.ShouldBind(&page); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		class_id := c.Request.URL.Query().Get("class_id")
		subject_id := c.Request.URL.Query().Get("subject_id")

		store := resultStorage.NewMongoStore(appCtx.GetNewDataMongoDB())
		biz := resultBiz.NewListMarkBiz(store)
		result, err := biz.ListResult(c.Request.Context(), bson.M{"class_id": class_id, "subject_id": subject_id}, &page)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse(result))
	}
}
