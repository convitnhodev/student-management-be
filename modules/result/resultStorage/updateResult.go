package resultStorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/result/resultModel"
)

func (db *mongoStore) UpdateResult(ctx context.Context, filter interface{}, data resultModel.Result) error {
	collection := db.db.Database(setupDatabase.NameDB).Collection(resultModel.NameCollection)

	update := bson.M{
		"$set": data,
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't Insert to DB, something DB is error")
		return solveError.ErrDB(err)
	}

	managerLog.InfoLogger.Println("Insert to DB success")
	return nil
}
