package notificationStorage

import (
	"context"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/notifedProvider/notificationModel"
)

func (db *mongoStore) CreateNewNotification(ctx context.Context, data *notificationModel.Notification) error {
	collection := db.db.Database(setupDatabase.NameDB).Collection(notificationModel.NameCollection)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't Insert Notification to DB, something DB is error")
		return solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("Insert Notification to DB success")
	return nil
}
