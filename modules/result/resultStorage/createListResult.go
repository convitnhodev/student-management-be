package resultStorage

import (
	"context"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/result/resultModel"
)

func (db *mongoStore) CreateListResult(ctx context.Context, data []resultModel.Result) error {
	collection := db.db.Database(setupDatabase.NameDB).Collection(resultModel.NameCollection)

	result := make([]interface{}, 0)
	for _, element := range data {
		result = append(result, element)
	}

	_, err := collection.InsertMany(ctx, result)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't Insert to DB, something DB is error")
		return solveError.ErrDB(err)
	}

	managerLog.InfoLogger.Println("Insert to DB success")
	return nil
}
