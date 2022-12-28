package courseStorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/course/courseModel"
)

func (db *mongoStore) FindCourse(ctx context.Context, conditions interface{}) (*courseModel.Course, error) {
	collection := db.db.Database(setupDatabase.NameDB).Collection(courseModel.NameCollection)

	var data bson.M

	if err := collection.FindOne(ctx, conditions).Decode(&data); err != nil {
		if err.Error() == solveError.RecordNotFound {
			managerLog.InfoLogger.Println("Cant find record from database")
			return nil, err
		}
		managerLog.ErrorLogger.Println("Can't find record into DB, something DB is error")
		return nil, solveError.ErrDB(err)
	}

	var result courseModel.Course
	bsonBytes, _ := bson.Marshal(data)
	bson.Unmarshal(bsonBytes, &result)
	managerLog.InfoLogger.Println("Find record success, storage return record and nil error")
	return &result, nil
}
