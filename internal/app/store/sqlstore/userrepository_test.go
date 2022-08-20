package sqlstore_test

import (
	"github.com/stretchr/testify/assert"
	"goRestApi/internal/app/model"
	"goRestApi/internal/app/store"
	"goRestApi/internal/app/store/sqlstore"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
	teardown("users")
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@test.test"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u1 := model.TestUser(t)
	u1.Email = email

	s.User().Create(u1)
	u2, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)

}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	s.User().Create(u)
	u, err := s.User().Find(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}


