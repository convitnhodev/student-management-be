package subjectStorage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"managerstudent/common/paging"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/component/managerLog"
	"managerstudent/modules/subject/subjectModel"
)

func (db *mongoStore) ListSubjects(ctx context.Context, filter interface{}, page *paging.Paging) ([]subjectModel.Subject, error) {
	collection := db.db.Database(setupDatabase.NameDB).Collection(subjectModel.NameCollection)

	opstions := new(options.FindOptions)
	opstions.SetLimit(page.Limit)
	opstions.SetSkip(int64(page.Page-1) * page.Limit)

	var result []subjectModel.Subject
	cursor, err := collection.Find(ctx, filter, opstions)
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
