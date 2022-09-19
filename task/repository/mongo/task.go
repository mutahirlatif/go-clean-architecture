package mongo

import (
	"context"
	"time"

	"github.com/mutahirlatif/go-clean-architecture/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Task struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserID     primitive.ObjectID `bson:"userId"`
	TaskDetail string             `bson:"taskDetail"`
	DueDate    time.Time          `bson:"dueDate"`
}

type TaskRepository struct {
	db *mongo.Collection
}

func NewTaskRepository(db *mongo.Database, collection string) *TaskRepository {
	return &TaskRepository{
		db: db.Collection(collection),
	}
}

func (r TaskRepository) CreateTask(ctx context.Context, user *models.User, tm *models.Task) error {
	tm.UserID = user.ID

	model := toModel(tm)

	res, err := r.db.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	tm.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r TaskRepository) GetTasks(ctx context.Context, user *models.User) ([]*models.Task, error) {
	uid, _ := primitive.ObjectIDFromHex(user.ID)
	cur, err := r.db.Find(ctx, bson.M{
		"userId": uid,
	})

	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	out := make([]*Task, 0)

	for cur.Next(ctx) {
		task := new(Task)
		err := cur.Decode(task)
		if err != nil {
			return nil, err
		}

		out = append(out, task)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return toTasks(out), nil
}

func (r TaskRepository) GetTaskByID(ctx context.Context, user *models.User, id string) (*models.Task, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	uID, _ := primitive.ObjectIDFromHex(user.ID)
	out := new(Task)
	err := r.db.FindOne(ctx, bson.M{"_id": objID, "userId": uID}).Decode(out)

	if err != nil {
		return nil, err
	}

	return toTask(out), nil
}

func (r TaskRepository) UpdateTask(ctx context.Context, user *models.User, tm *models.Task, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	uID, _ := primitive.ObjectIDFromHex(user.ID)
	tm.UserID = user.ID
	model := toModel(tm)

	_, err := r.db.ReplaceOne(ctx, bson.M{"_id": objID, "userId": uID}, model)
	if err != nil {
		return err
	}

	return nil
}

func (r TaskRepository) DeleteTask(ctx context.Context, user *models.User, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	uID, _ := primitive.ObjectIDFromHex(user.ID)

	_, err := r.db.DeleteOne(ctx, bson.M{"_id": objID, "userId": uID})
	return err
}

func toModel(t *models.Task) *Task {
	uid, _ := primitive.ObjectIDFromHex(t.UserID)

	return &Task{
		UserID:     uid,
		DueDate:    t.DueDate,
		TaskDetail: t.TaskDetail,
	}
}

func toTask(t *Task) *models.Task {
	return &models.Task{
		ID:         t.ID.Hex(),
		UserID:     t.UserID.Hex(),
		TaskDetail: t.TaskDetail,
		DueDate:    t.DueDate,
	}
}

func toTasks(bs []*Task) []*models.Task {
	out := make([]*models.Task, len(bs))

	for i, b := range bs {
		out[i] = toTask(b)
	}

	return out
}
