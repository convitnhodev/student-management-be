package courseStorage

import (
	"context"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
)

func (db *mongoStore) DeleteCourse(ctx context.Context, conditions interface{}) error {
	collection := db.db.Database("ManagerStudent").Collection("Course")
	if _, err := collection.DeleteOne(ctx, conditions); err != nil {
		managerLog.ErrorLogger.Println("something DB is error")
		return solveError.ErrDB(err)

	}
	return nil
}
