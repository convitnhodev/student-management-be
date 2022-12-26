package notifyStorage

import (
	"context"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/notifedProvider/notifyModel"
)

func (db *mongoStore) CreateNewNotify(ctx context.Context, data *notifyModel.Notify) error {
	collection := db.db.Database("ManagerStudent").Collection("Notify")
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't Insert to DB, something DB is error")
		return solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("Insert to DB success")
	return nil
}
