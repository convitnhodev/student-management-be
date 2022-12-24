package markStorage

import (
	"context"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/student/studentModel"
)

func (db *mongoStore) ListResultByConditions(ctx context.Context, conditions interface{}) (*studentModel.Result, error) {
	collection := db.db.Database("ManagerStudent").Collection("Mark")

	dataCursor, err := collection.Find(ctx, conditions)
	if err != nil {
		if err.Error() == solveError.RecordNotFound {
			managerLog.InfoLogger.Println("Cant find record from database")
			return nil, err
		}
		managerLog.ErrorLogger.Println("Can't find record into DB, something DB is error")
		return nil, solveError.ErrDB(err)
	}

	var result studentModel.Result
	if err := dataCursor.All(ctx, &result); err != nil {
		managerLog.ErrorLogger.Println("Can't find record into DB, something DB is error")
		return nil, solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("Find record success, storage return record and nil error")
	return &result, nil
}
