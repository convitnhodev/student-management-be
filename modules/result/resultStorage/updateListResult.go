package resultStorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/result/resultModel"
)

func (db *mongoStore) UpdateResult(ctx context.Context, conditions interface{}, data resultModel.Result) error {
	collection := db.db.Database(setupDatabase.NameDB).Collection(resultModel.NameCollection)
	tmp, _ := bson.Marshal(data)
	var target bson.D

	_ = bson.Unmarshal(tmp, &target)
	update := bson.M{
		"$set": target,
	}
	_, err := collection.UpdateOne(ctx, conditions, update)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't update to DB, something DB is error")
		return solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("update result success, storage return record and nil error")
	return nil
}
