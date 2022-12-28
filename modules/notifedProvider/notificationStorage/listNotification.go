package notificationStorage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"managerstudent/common/paging"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/notifedProvider/notificationModel"
)

func (db *mongoStore) ListNotificationsByConditions(ctx context.Context, conditions interface{}, page *paging.Paging) ([]notificationModel.Notification, error) {
	collection := db.db.Database(setupDatabase.NameDB).Collection(notificationModel.NameCollection)

	opstions := new(options.FindOptions)
	opstions.SetLimit(page.Limit)
	opstions.SetSkip(int64(page.Page-1) * page.Limit)
	page.Total, _ = collection.CountDocuments(ctx, conditions)
	dataCursor, err := collection.Find(ctx, conditions, opstions)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't find record into DB, something DB is error")
		return nil, solveError.ErrDB(err)
	}

	var notifications []notificationModel.Notification

	if err := dataCursor.All(ctx, &notifications); err != nil {
		managerLog.ErrorLogger.Println("Can't find record into DB, something DB is error")
		return nil, solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("Find record success, storage return record and nil error")

	return notifications, nil
}
