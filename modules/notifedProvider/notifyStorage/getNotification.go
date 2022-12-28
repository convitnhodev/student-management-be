package notifyStorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/notifedProvider/notifyModel"
)

func (db *mongoStore) GetNotification(ctx context.Context, conditions interface{}) (*notifyModel.Notification, error) {
	collection := db.db.Database("ManagerStudent").Collection("Notification")

	var data bson.M

	err := collection.FindOne(ctx, conditions).Decode(&data)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't find notification from DB, something DB is error")
		return nil, solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("get info from DB success")
	var result notifyModel.Notification
	bsonBytes, _ := bson.Marshal(data)
	bson.Unmarshal(bsonBytes, &result)
	managerLog.InfoLogger.Println("Find record success, storage return record and nil error")
	return &result, nil
}
