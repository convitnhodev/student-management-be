package subjectTransport

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/customResponse"
	"managerstudent/component"
	"managerstudent/modules/subject/subjectBiz"
	"managerstudent/modules/subject/subjectStorage"
)

func DeleteSubject(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := c.GetQuery("id")
		if ok == false {
			panic("id is not exist")
		}
		store := subjectStorage.NewMongoStore(appCtx.GetNewDataMongoDB())
		biz := subjectBiz.NewDeleteSubjectBiz(store)
		if err := biz.DeleteSubject(c.Request.Context(), bson.M{"id": id}); err != nil {
			c.JSON(400, err)
			return
		}

		c.JSON(200, customResponse.SimpleSuccessReponse("success"))
	}
}
