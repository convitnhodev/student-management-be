package userStorage

import (
	"context"
	"managerstudent/common/paging"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/user/userModel"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *mongoStore) ListUsersByConditions(ctx context.Context, conditions interface{}, page *paging.Paging) ([]userModel.User, error) {
	collection := db.db.Database(setupDatabase.NameDB).Collection(userModel.NameCollection)

	options := new(options.FindOptions)
	options.SetLimit(page.Limit)
	options.SetSkip(int64(page.Page-1) * page.Limit)
	page.Total, _ = collection.CountDocuments(ctx, conditions)
	dataCursor, err := collection.Find(ctx, conditions, options)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't find record into DB, something DB is error")
		return nil, solveError.ErrDB(err)
	}

	var users []userModel.User

	if err := dataCursor.All(ctx, &users); err != nil {
		managerLog.ErrorLogger.Println("Can't find record into DB, something DB is error")
		return nil, solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("Find record success, storage return record and nil error")

	return users, nil
}
