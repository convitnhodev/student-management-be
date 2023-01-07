package studentStorage

import (
	"context"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/student/studentModel"

	"go.mongodb.org/mongo-driver/bson"
)

func (db *mongoStore) UpdateStudent(ctx context.Context, conditions interface{}, student studentModel.Student) error {
	collection := db.db.Database(setupDatabase.NameDB).Collection(studentModel.NameCollection)

	data := bson.M{
		"$set": student,
	}

	_, err := collection.UpdateOne(ctx, conditions, data)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't update to DB, something DB is error")
		return solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("update result success, storage return record and nil error")
	return nil
}
