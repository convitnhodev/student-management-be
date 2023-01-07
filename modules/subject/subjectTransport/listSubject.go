package subjectTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/customResponse"
	"managerstudent/component"
	"managerstudent/modules/subject/subjectBiz"
	"managerstudent/modules/subject/subjectStorage"
)



func ListSubjects(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		filter, _ := c.GetQuery("class_id")
		store := subjectStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := subjectBiz.NewListSubjectBiz(store)
		data, err := biz.ListSubject(c.Request.Context(), filter)
		if err != nil {
			c.JSON(400, err)
			return
		}

		c.JSON(200, customResponse.SimpleSuccessReponse(data))

	}
}