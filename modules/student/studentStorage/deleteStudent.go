package studentStorage

import (
	"context"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
)

func (db *mongoStore) DeleteStudent(ctx context.Context, filter interface{}, location string) error {
	collection := db.db.Database("ManagerStudent").Collection(location)
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't delete to DB, something DB is error")
		return solveError.ErrDB(err)
	}

	managerLog.InfoLogger.Println("delete to DB success")
	return nil
}
