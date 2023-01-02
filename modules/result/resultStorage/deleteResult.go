package resultStorage

import (
	"context"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/result/resultModel"
)

func (db *mongoStore) DeleteResult(ctx context.Context, filter interface{}) error {
	collection := db.db.Database(setupDatabase.NameDB).Collection(resultModel.NameCollection)

	_, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't Delete into DB, something DB is error")
		return solveError.ErrDB(err)
	}

	managerLog.InfoLogger.Println("Delete into DB success")
	return nil
}
