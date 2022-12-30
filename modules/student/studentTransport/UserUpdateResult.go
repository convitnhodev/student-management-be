package studentTransport

//func UserUpdateResult(app component.AppContext) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var data studentModel.UpdateResult
//		if err := c.ShouldBind(&data); err != nil {
//			panic(solveError.ErrInvalidRequest(err))
//		}
//		store := studentStorage.NewMongoStore(app.GetNewDataMongoDB())
//		biz := studentBiz.NewUserUpdateResultStoreBiz(store, app.GetPubsub())
//		err := biz.UserUpdateResult(c.Request.Context(), &data)
//		if err != nil {
//			c.JSON(400, err)
//			return
//		}
//		c.JSON(200, customResponse.SimpleSuccessReponse(data.StudentId))
//	}
//}
