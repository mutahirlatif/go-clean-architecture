package postgres

import (
	"context"
	"strconv"

	"github.com/mutahirlatif/go-clean-architecture/models"
	"gorm.io/gorm"
)

type User struct {
	ID       int    `gorm:"primary_key;auto_increment"`
	Username string `gorm:"size:255;not null;unique"`
	Password string `gorm:"size:100;not null;"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB, collection string) *UserRepository {
	user := User{}
	db.AutoMigrate(&user)
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	model := toPostGresUser(user)

	var err = r.db.Debug().Create(&model).Error
	if err != nil {
		return err
	}
	return nil
}

func (r UserRepository) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	user := new(User)
	var err = r.db.Debug().Model(User{}).Where("username = ? AND password = ?", username, password).Take(user).Error
	if err != nil {
		return toModel(&User{}), err
	}
	return toModel(user), nil
}

func toPostGresUser(u *models.User) *User {
	return &User{
		Username: u.Username,
		Password: u.Password,
	}
}

func toModel(u *User) *models.User {
	id := strconv.Itoa(u.ID)
	return &models.User{
		ID:       id,
		Username: u.Username,
		Password: u.Password,
	}
}
