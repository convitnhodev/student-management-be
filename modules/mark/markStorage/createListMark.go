package markStorage

import (
	"context"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/mark/markModel"
)

func (db *mongoStore) CreateListMark(ctx context.Context, data []markModel.Result) error {
	collection := db.db.Database("ManagerStudent").Collection("Result")

	result := make([]interface{}, 0)
	for element := range data {
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
