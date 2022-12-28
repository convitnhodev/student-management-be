package courseStorage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"managerstudent/common/paging"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/course/courseModel"
)

func (db *mongoStore) ListCourses(ctx context.Context, conditions interface{}, page *paging.Paging) ([]courseModel.Course, error) {
	collection := db.db.Database(setupDatabase.NameDB).Collection(courseModel.NameCollection)

	opstions := new(options.FindOptions)
	opstions.SetLimit(page.Limit)
	opstions.SetSkip(int64(page.Page-1) * page.Limit)

	dataCursor, err := collection.Find(ctx, conditions, opstions)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't find record into DB, something DB is error")
		return nil, solveError.ErrDB(err)
	}

	var result []courseModel.Course
	if err := dataCursor.All(ctx, &result); err != nil {
		managerLog.ErrorLogger.Println("Can't find record into DB, something DB is error")
		return nil, solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("Find record success, storage return record and nil error")
	return result, nil
}
