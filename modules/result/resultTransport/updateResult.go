package resultTransport

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/customResponse"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/result/resultBiz"
	"managerstudent/modules/result/resultModel"
	"managerstudent/modules/result/resultStorage"
)

func UpdateResult(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result resultModel.Result
		if err := c.ShouldBind(&result); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		filter := bson.D{{"student_id", result.StudentId}, {"course_id", result.CourseId}}

		store := resultStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := resultBiz.NewUpdateResultBiz(store)
		if err := biz.UpdateResult(c.Request.Context(), filter, result); err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse("success"))
	}
}
