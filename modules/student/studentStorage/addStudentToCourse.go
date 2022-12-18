package studentStorage

import (
	"context"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/student/studentModel"
)

func (db *mongoStore) CreateNewStudentInCourse(ctx context.Context, data []studentModel.Student) error {
	collection := db.db.Database("ManagerStudent").Collection("User_Course")
	newStudent := make([]interface{}, len(data))
	for i := 0; i < len(data); i++ {
		newStudent[i] = data[i]
	}

	_, err := collection.InsertMany(ctx, newStudent)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't Insert to DB, something DB is error")
		return solveError.ErrDB(err)
	}

	managerLog.InfoLogger.Println("Insert to DB success")
	return nil
}
