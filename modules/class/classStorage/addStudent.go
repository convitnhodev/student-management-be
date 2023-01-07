package classStorage

import (
	"context"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/class/classModel"

	"go.mongodb.org/mongo-driver/bson"
)

func (db *mongoStore) AddStudent(ctx context.Context, condition interface{}, studentId string) error {
	collection := db.db.Database(setupDatabase.NameDB).Collection(classModel.NameCollection)

	data := bson.M{
		"$push": bson.M{
			"list_student_id": studentId,
		},
	}

	_, err := collection.UpdateOne(ctx, condition, data)
	if err != nil {
		managerLog.ErrorLogger.Println("Can't update to DB, something DB is error")
		return solveError.ErrDB(err)
	}
	managerLog.InfoLogger.Println("update result success, storage return record and nil error")
	return nil
}
