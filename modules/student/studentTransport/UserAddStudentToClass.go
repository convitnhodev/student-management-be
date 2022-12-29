package studentTransport

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"managerstudent/common/customResponse"
	"managerstudent/common/solveError"
	"managerstudent/component"
	"managerstudent/modules/student/studentBiz"
	"managerstudent/modules/student/studentModel"
	"managerstudent/modules/student/studentStorage"
)

func UserAddStudentToClass(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data studentModel.StudentAndClass
		fmt.Println(data)
		if err := c.ShouldBind(&data); err != nil {
			panic(solveError.ErrInvalidRequest(err))
		}
		store := studentStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := studentBiz.NewUserCreateStudentToClassBiz(store, app.GetPubsub())
		err := biz.UserCreateNewStudentToClass(c.Request.Context(), &data)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse(data.StudentId))
	}
}
