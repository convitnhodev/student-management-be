package subjectStorage

import (
	"context"
	"managerstudent/common/setupDatabase"
	"managerstudent/common/solveError"
	"managerstudent/modules/subject/subjectModel"
)

func (db *mongoStore) CreateNewSubject(ctx context.Context, data *subjectModel.Subject) error {
	collection := db.db.Database(setupDatabase.NameDB).Collection(subjectModel.NameCollection)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return solveError.ErrDB(err)
	}
	return nil
}