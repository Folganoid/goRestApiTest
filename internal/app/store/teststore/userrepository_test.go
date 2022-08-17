package teststore_test

import (
	"github.com/stretchr/testify/assert"
	"goRestApi/internal/app/model"
	"goRestApi/internal/app/store"
	"goRestApi/internal/app/store/teststore"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {

	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {

	s := teststore.New()
	email := "user@test.test"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(model.TestUser(t))
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}



