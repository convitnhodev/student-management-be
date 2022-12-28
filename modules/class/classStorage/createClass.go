package classStorage

import (
	"context"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/class/classModel"
)

func (db *mongoStore) CreateNewClass(ctx context.Context, data *classModel.Class) error {
	collection := db.db.Database(setupDatabase.NameDB).Collection(classModel.NameCollection)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't Insert New Class to DB, something DB is error")
		return solveError.ErrDB(err)
	}

	managerLog.InfoLogger.Println("Insert New Class to DB success")
	return nil
}
