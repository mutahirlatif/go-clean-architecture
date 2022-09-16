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
	ID         int `gorm:"primary_key;auto_increment;<-:create"`
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
	uID, _ := strconv.Atoi(user.ID)
	tasks := []Task{}
	var err = r.db.Debug().Model(Task{}).Where("user_id = ?", uID).Find(&tasks).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return toTasks(tasks), nil

}

func (r TaskRepository) GetTaskByID(ctx context.Context, user *models.User, id string) (*models.Task, error) {
	objID, _ := strconv.Atoi(id)
	uID, _ := strconv.Atoi(user.ID)

	task := Task{}
	var err = r.db.Debug().Model(Task{}).Where("user_id = ? AND id = ?", uID, objID).First(&task).Error
	if err != nil {
		return nil, err
	}

	return toTask(&task), nil

}

func (r TaskRepository) UpdateTask(ctx context.Context, user *models.User, tm *models.Task, id string) error {
	objID, _ := strconv.Atoi(id)
	uID, _ := strconv.Atoi(user.ID)

	task := Task{}
	var err = r.db.Debug().Model(Task{}).Where("user_id = ? AND id = ?", uID, objID).First(&task).Error
	if err != nil {
		return err
	}

	task.TaskDetail = tm.TaskDetail
	task.DueDate = tm.DueDate

	return r.db.Debug().Model(&task).Save(&task).Error
}

func (r TaskRepository) DeleteTask(ctx context.Context, user *models.User, id string) error {
	objID, _ := strconv.Atoi(id)
	uID, _ := strconv.Atoi(user.ID)

	task := Task{}
	var err = r.db.Debug().Model(Task{}).Where("user_id = ? AND id = ?", uID, objID).First(&task).Error
	if err != nil {
		return err
	}

	r.db.Debug().Model(&task).Delete(&task)
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

func toTasks(ts []Task) []*models.Task {
	out := make([]*models.Task, len(ts))

	for i, t := range ts {
		out[i] = toTask(&t)
	}
	return out
}
