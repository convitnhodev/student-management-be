package notifyStorage

import (
	"context"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/notifedProvider/notifyModel"
)

func (db *mongoStore) CreateNewNotification(ctx context.Context, data *notifyModel.Notification) error {
	collection := db.db.Database("ManagerStudent").Collection("Notification")
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't Insert Notification to DB, something DB is error")
		return solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("Insert Notification to DB success")
	return nil
}
