package studentStorage

import (
	"context"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/student/studentModel"
)

func (db *mongoStore) CreateNewStudent(ctx context.Context, data studentModel.Student) error {
	collection := db.db.Database(setupDatabase.NameDB).Collection(studentModel.NameCollection)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't Insert to DB, something DB is error")
		return solveError.ErrDB(err)
	}

	managerLog.InfoLogger.Println("Insert to DB success")
	return nil
}
