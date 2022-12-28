package notificationStorage

import (
	"context"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/notifedProvider/notificationModel"
)

func (db *mongoStore) SolveNotification(ctx context.Context, conditions interface{}, value interface{}) error {
	collection := db.db.Database(setupDatabase.NameDB).Collection(notificationModel.NameCollection)
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
