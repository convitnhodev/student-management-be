package notifyStorage

import (
	"context"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/notifedProvider/notifyModel"
)

func (db *mongoStore) SolveNotify(ctx context.Context, data *notifyModel.Notify, conditions interface{}, value interface{}) error {
	collection := db.db.Database("ManagerStudent").Collection("Notify")
	//filter := bson.D{{"id", data.Id}}
	//update := bson.D{{"$set", bson.D{{"status", value}}}}

	_, err := collection.UpdateOne(ctx, conditions, value)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't update to DB, something DB is error")
		return solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("update result success, storage return record and nil error")
	return nil
}
