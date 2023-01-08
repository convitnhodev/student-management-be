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

func GetEachResult(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var page paging.Paging
		if err := c.ShouldBind(&page); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		student_id := c.Request.URL.Query().Get("student_id")
		subject_id := c.Request.URL.Query().Get("subject_id")

		store := resultStorage.NewMongoStore(appCtx.GetNewDataMongoDB())
		biz := resultBiz.NewGetMarkBiz(store)
		result, err := biz.GetEachResult(c.Request.Context(), bson.M{"student_id": student_id, "subject_id": subject_id})
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse(result))
	}
}
