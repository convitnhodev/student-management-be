package subjectStorage

import (
	"context"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/subject/subjectModel"
)

func (db *mongoStore) ListSubjects(ctx context.Context, filter interface{}) ([]subjectModel.Subject, error) {
	collection := db.db.Database(setupDatabase.NameDB).Collection(subjectModel.NameCollection)
	var result []subjectModel.Subject
	cursor, err := collection.Find(ctx,filter)
	if err != nil {
		if err.Error() == solveError.RecordNotFound {
			managerLog.InfoLogger.Println("Cant find record from database")
			return nil, err
		}
		managerLog.ErrorLogger.Println("Can't find record into DB, something DB is error")
		return nil, solveError.ErrDB(err)
	}
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}
