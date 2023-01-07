package subjectTransport

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/customResponse"
	"managerstudent/component"
	"managerstudent/modules/subject/subjectBiz"
	"managerstudent/modules/subject/subjectStorage"
)

func GetSubject(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, ok := c.GetQuery("id")
		if ok == false {
			panic("id is not exist")
		}

		store := subjectStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := subjectBiz.NewGetSubjectBiz(store)
		data, err := biz.GetSubject(c.Request.Context(), bson.M{"id": id})
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, customResponse.SimpleSuccessReponse(data))

	}
}
