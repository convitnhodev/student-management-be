package markStorage

import (
	"context"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/student/studentModel"
)

func (db *mongoStore) UpdateResult(ctx context.Context, data []studentModel.Result, conditions interface{}) error {
	collection := db.db.Database("ManagerStudent").Collection("Mark")

	_, err := collection.UpdateMany(ctx, conditions, data)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't update to DB, something DB is error")
		return solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("update result success, storage return record and nil error")
	return nil
}
