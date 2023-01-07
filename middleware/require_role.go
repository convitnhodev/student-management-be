package middleware

import (
	"github.com/gin-gonic/gin"
	_const "managerstudent/common/const"
	"managerstudent/common/solveError"
	"managerstudent/component"
)

func RequireRole(appCtx component.AppContext, role ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(_const.CurrentUser).(_const.Requester)
		for i := range role {
			if u.GetRoleInt() == _const.Role((role[i])) {
				c.Next()
				return
			}
		}

		panic(solveError.ErrNoPermission(nil))

	}
}
