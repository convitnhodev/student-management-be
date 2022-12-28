package courseStorage

import (
	"context"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/course/courseModel"
)

func (db *mongoStore) DeleteCourse(ctx context.Context, conditions interface{}) error {
	collection := db.db.Database(setupDatabase.NameDB).Collection(courseModel.NameCollection)
	if _, err := collection.DeleteOne(ctx, conditions); err != nil {
		managerLog.ErrorLogger.Println("something DB is error")
		return solveError.ErrDB(err)

	}
	return nil
}
