package classStorage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"managerstudent/common/paging"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/class/classModel"
)

func (db *mongoStore) ListClasses(ctx context.Context, conditions interface{}, page *paging.Paging) ([]classModel.Class, error) {
	collection := db.db.Database("ManagerStudent").Collection("Class")

	opstions := new(options.FindOptions)
	opstions.SetLimit(page.Limit)
	opstions.SetSkip(int64(page.Page-1) * page.Limit)

	dataCursor, err := collection.Find(ctx, conditions, opstions)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't find record into DB, something DB is error")
		return nil, solveError.ErrDB(err)
	}

	var result []classModel.Class
	if err := dataCursor.All(ctx, &result); err != nil {
		managerLog.ErrorLogger.Println("Can't find record into DB, something DB is error")
		return nil, solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("Find record success, storage return record and nil error")
	return result, nil
}
