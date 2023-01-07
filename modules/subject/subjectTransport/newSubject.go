package subjectTransport

import (
	"github.com/gin-gonic/gin"
	"managerstudent/common/customResponse"
	"managerstudent/component"
	"managerstudent/modules/subject/subjectBiz"
	"managerstudent/modules/subject/subjectModel"
	"managerstudent/modules/subject/subjectStorage"
)

func NewCreateSubject (app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var subject subjectModel.Subject
		if err := c.ShouldBind(&subject); err != nil {
			c.JSON(400, err)
			return
		}
		store := subjectStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := subjectBiz.NewCreateSubjectBiz(store)
		if err := biz.CreateSubject(c.Request.Context(), &subject); err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse(subject))
	}
}
