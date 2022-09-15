package postgres

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/mutahirlatif/go-clean-architecture/models"
	"gorm.io/gorm"
)

type Task struct {
	ID         int `gorm:"primary_key;auto_increment"`
	UserID     int
	TaskDetail string `gorm:"size:255;"`
	DueDate    time.Time
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB, collection string) *TaskRepository {
	task := Task{}
	db.AutoMigrate(&task)
	return &TaskRepository{
		db: db,
	}
}

func (r TaskRepository) CreateTask(ctx context.Context, user *models.User, tm *models.Task) error {
	model := toPostGresTask(user, tm)
	var err = r.db.Debug().Create(&model).Error
	if err != nil {
		return err
	}
	return nil

}

func (r TaskRepository) GetTasks(ctx context.Context, user *models.User) ([]*models.Task, error) {
	//Task := new(Task)
	uid, _ := strconv.Atoi(user.ID)
	tasks := []Task{}
	var err = r.db.Debug().Model(Task{}).Where("user_id = ?", uid).First(&tasks).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	out := make([]*Task, 0)

	for _, t := range tasks {
		out = append(out, &t)
	}
	return toTasks(out), nil

}

func (r TaskRepository) GetTaskByID(ctx context.Context, user *models.User, id string) (*models.Task, error) {
	// objID, _ := primitive.ObjectIDFromHex(id)
	// uID, _ := primitive.ObjectIDFromHex(user.ID)
	// out := new(Task)
	// err := r.db.FindOne(ctx, bson.M{"_id": objID, "userId": uID}).Decode(out)

	// if err != nil {
	// 	return nil, err
	// }

	// return toTask(out), nil
	return nil, nil
}

func (r TaskRepository) UpdateTask(ctx context.Context, user *models.User, tm *models.Task, id string) error {
	// objID, _ := primitive.ObjectIDFromHex(id)
	// uID, _ := primitive.ObjectIDFromHex(user.ID)
	// tm.UserID = user.ID
	// model := toModel(tm)

	// _, err := r.db.ReplaceOne(ctx, bson.M{"_id": objID, "userId": uID}, model)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (r TaskRepository) DeleteTask(ctx context.Context, user *models.User, id string) error {
	// objID, _ := primitive.ObjectIDFromHex(id)
	// uID, _ := primitive.ObjectIDFromHex(user.ID)

	// _, err := r.db.DeleteOne(ctx, bson.M{"_id": objID, "userId": uID})
	// return err
	return nil
}

func toPostGresTask(user *models.User, t *models.Task) *Task {
	// [TODO] Check error
	uid, _ := strconv.Atoi(user.ID)

	return &Task{
		UserID:     uid,
		DueDate:    t.DueDate,
		TaskDetail: t.TaskDetail,
	}
}

func toTask(t *Task) *models.Task {
	id := strconv.Itoa(t.ID)
	uid := strconv.Itoa(t.UserID)
	return &models.Task{
		ID:         id,
		UserID:     uid,
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
