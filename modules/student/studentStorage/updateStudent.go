package studentStorage

import (
	"context"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
)

func (db *mongoStore) UpdateStudent(ctx context.Context, conditions interface{}, value interface{}, location string) error {
	collection := db.db.Database("ManagerStudent").Collection(location)

	_, err := collection.UpdateOne(ctx, conditions, value)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't update to DB, something DB is error")
		return solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("update result success, storage return record and nil error")
	return nil
}
