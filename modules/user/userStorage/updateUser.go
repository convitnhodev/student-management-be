package userStorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/user/userModel"
)

func (db *mongoStore) UpdateUser(ctx context.Context, conditions interface{}, data interface{}) error {
	collection := db.db.Database(setupDatabase.NameDB).Collection(userModel.NameCollection)
	update := bson.M{
		"$set": data,
	}
	_, err := collection.UpdateOne(ctx, conditions, update)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't update to DB, something DB is error")
		return solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("update result success, storage return record and nil error")
	return nil
}
