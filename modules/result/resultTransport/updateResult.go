package resultTransport

import (
	"managerstudent/common/customResponse"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/result/resultBiz"
	"managerstudent/modules/result/resultModel"
	"managerstudent/modules/result/resultStorage"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateResult(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var results []resultModel.Result
		if err := c.ShouldBind(&results); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		for _, result := range results {
			filter := bson.D{{"student_id", result.StudentId}, {"subject_id", result.SubjectId}}
			store := resultStorage.NewMongoStore(app.GetNewDataMongoDB())
			biz := resultBiz.NewUpdateResultBiz(store)
			if err := biz.UpdateResult(c.Request.Context(), filter, result); err != nil {
				c.JSON(400, err)
				return
			}
		}

		c.JSON(200, customResponse.SimpleSuccessReponse("success"))
	}
}
