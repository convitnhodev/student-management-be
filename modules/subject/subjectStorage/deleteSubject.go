package subjectStorage

import (
	"context"
	"managerstudent/common/setupDatabase"
	"managerstudent/modules/subject/subjectModel"
)

func (db *mongoStore) DeleteSubject(ctx context.Context, conditions interface{}) error {
	collection := db.db.Database(setupDatabase.NameDB).Collection(subjectModel.NameCollection)
	if _, err := collection.DeleteOne(ctx, conditions); err != nil {
		return err
	}
	return nil
}
