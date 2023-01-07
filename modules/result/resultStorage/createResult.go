package resultStorage

import (
	"context"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/result/resultModel"
)

func (db *mongoStore) CreateResult(ctx context.Context, data []resultModel.Result) error {
	collection := db.db.Database(setupDatabase.NameDB).Collection(resultModel.NameCollection)

	newData := make([]interface{}, len(data))

	for i := range data {
		newData[i] = data[i]
	}
	_, err := collection.InsertMany(ctx, newData)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't Insert to DB, something DB is error")
		return solveError.ErrDB(err)
	}

	managerLog.InfoLogger.Println("Insert to DB success")
	return nil
}
