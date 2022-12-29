package studentTransport

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"managerstudent/common/solveError"
	"managerstudent/component"
)

func PostAgentLogFile(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		fileHeader, err := c.FormFile("file")

		if err != nil {
			solveError.ErrInvalidRequest(err)
		}

		c.SaveUploadedFile(fileHeader, fmt.Sprintf("./filesave/%s", fileHeader.Filename))
		if err != nil {
			solveError.ErrInvalidRequest(err)
		}

		c.JSON(200, "SUCCESS")
	}
}
