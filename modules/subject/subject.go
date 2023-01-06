package subject

import (
	"context"
	"fmt"
	"managerstudent/common/setupDatabase"
	"managerstudent/component"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const NameCollection = "Subjects"

type Subject struct {
	ID    string `json:"id" bson:"_id"`
	Title string `json:"title" bson:"title"`
}
type mongoStore struct {
	db *mongo.Client
}

func NewMongoStore(db *mongo.Client) *mongoStore {
	return &mongoStore{db}
}

func (db *mongoStore) CreateNewSubject(ctx context.Context, data *Subject) error {
	collection := db.db.Database(setupDatabase.NameDB).Collection(NameCollection)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (db *mongoStore) ListSubjects(ctx context.Context) ([]Subject, error) {
	collection := db.db.Database(setupDatabase.NameDB).Collection(NameCollection)
	var result []Subject
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (db *mongoStore) DeleteSubject(ctx context.Context, conditions interface{}) error {
	collection := db.db.Database(setupDatabase.NameDB).Collection(NameCollection)
	if _, err := collection.DeleteOne(ctx, conditions); err != nil {
		return err
	}
	return nil
}

func CreateSubject(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data Subject
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		if err := NewMongoStore(app.GetNewDataMongoDB()).CreateNewSubject(c, &data); err != nil {
			panic(err)
		}
		c.JSON(200, "Create subject success")
	}
}

func ListSubjects(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := NewMongoStore(app.GetNewDataMongoDB()).ListSubjects(c)
		if err != nil {
			panic(err)
		}
		c.JSON(200, result)
	}
}

func DeleteSubject(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Request.URL.Query().Get("id")
		fmt.Println(id)
		if err := NewMongoStore(app.GetNewDataMongoDB()).DeleteSubject(c, bson.M{"_id": id}); err != nil {
			panic(err)
		}
		c.JSON(200, "Delete subject success")
	}
}
