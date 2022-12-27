package notifyStorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/notifedProvider/notifyModel"
)

func (db *mongoStore) GetNotify(ctx context.Context, conditions interface{}) (*notifyModel.Notify, error) {
	collection := db.db.Database("ManagerStudent").Collection("Notify")

	var data bson.M

	err := collection.FindOne(ctx, conditions).Decode(&data)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't Insert to DB, something DB is error")
		return nil, solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("Insert to DB success")
	var result notifyModel.Notify
	bsonBytes, _ := bson.Marshal(data)
	bson.Unmarshal(bsonBytes, &result)
	managerLog.InfoLogger.Println("Find record success, storage return record and nil error")
	return &result, nil
}
