package localstorage

import (
	"context"
	"testing"

	"github.com/mutahirlatif/go-clean-architecture/auth"
	"github.com/mutahirlatif/go-clean-architecture/models"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	s := NewUserLocalStorage()

	id1 := "id"

	user := &models.User{
		ID:       id1,
		Username: "user",
		Password: "password",
	}

	err := s.CreateUser(context.Background(), user)
	assert.NoError(t, err)

	returnedUser, err := s.GetUser(context.Background(), "user", "password")
	assert.NoError(t, err)
	assert.Equal(t, user, returnedUser)

	_, err = s.GetUser(context.Background(), "user", "")
	assert.Error(t, err)
	assert.Equal(t, err, auth.ErrUserNotFound)
}
