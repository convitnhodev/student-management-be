package userStorage

import (
	"context"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/user/userModel"
)

func (db *mongoStore) CreateUser(ctx context.Context, data *userModel.User) error {
	collection := db.db.Database(setupDatabase.NameDB).Collection(userModel.NameCollection)

	if _, err := collection.InsertOne(ctx, data); err != nil {
		managerLog.ErrorLogger.Println("Can't Insert to DB, something DB is error")
		return solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("Create record success, return nil error")
	return nil
}
