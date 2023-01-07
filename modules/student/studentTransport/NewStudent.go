package studentTransport

import (
	"managerstudent/common/customResponse"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/student/studentModel"
	"managerstudent/modules/student/studentStorage"

	"github.com/gin-gonic/gin"
)

func AddStudent(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var student studentModel.Student
		if err := c.ShouldBind(&student); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}
		store := studentStorage.NewMongoStore(app.GetNewDataMongoDB())
		err := store.CreateNewStudent(c.Request.Context(), student)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse("success"))
	}
}
