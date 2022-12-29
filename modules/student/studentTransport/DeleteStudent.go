package studentTransport

//func DeleteStudent(app component.AppContext) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var id string
//		var ok bool
//
//		id, ok = c.GetQuery("id")
//		if !ok {
//			panic(solveError.ErrInvalidRequest(nil))
//		}
//		store := studentStorage.NewMongoStore(app.GetNewDataMongoDB())
//		biz := studentBiz.NewDeleteStudent(store)
//		if err := biz.DeleteStudent(c.Request.Context(), &id); err != nil {
//			c.JSON(400, err)
//			return
//		}
//		c.JSON(200, "ok")
//	}
//}
//
//func DeleteStudentFromClass(app component.AppContext) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var id string
//		var ok bool
//
//		id, ok = c.GetQuery("id")
//		if !ok {
//			panic(solveError.ErrInvalidRequest(nil))
//		}
//		store := studentStorage.NewMongoStore(app.GetNewDataMongoDB())
//		biz := studentBiz.NewDeleteStudent(store)
//		if err := biz.DeleteStudentFromClass(c.Request.Context(), &id); err != nil {
//			c.JSON(400, err)
//			return
//		}
//		c.JSON(200, "ok")
//	}
//}
//
//func DeleteStudentFromCourse(app component.AppContext) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var id string
//		var ok bool
//
//		id, ok = c.GetQuery("id")
//		if !ok {
//			panic(solveError.ErrInvalidRequest(nil))
//		}
//		store := studentStorage.NewMongoStore(app.GetNewDataMongoDB())
//		biz := studentBiz.NewDeleteStudent(store)
//		if err := biz.DeleteStudentFromCourse(c.Request.Context(), &id); err != nil {
//			c.JSON(400, err)
//			return
//		}
//		c.JSON(200, "ok")
//	}
//}
