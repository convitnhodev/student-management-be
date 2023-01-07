package resultTransport

import (
	"managerstudent/common/customResponse"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/result/resultBiz"
	"managerstudent/modules/result/resultModel"
	"managerstudent/modules/result/resultStorage"
	"managerstudent/modules/student/studentStorage"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func AddResult(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var results []resultModel.Result
		if err := c.ShouldBind(&results); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}

		for i := range results {
			results[i].CalculateAverage()
			studentStore := studentStorage.NewMongoStore(app.GetNewDataMongoDB())
			student, _ := studentStore.FindStudent(
				c.Request.Context(),
				bson.D{{"id", results[i].StudentId}},
			)
			results[i].ClassId = student.ClassId
		}

		store := resultStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := resultBiz.NewAddResultBiz(store)
		if err := biz.CreateOrUpdateResult(c.Request.Context(), results); err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse("success"))
	}
}
