package model

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserModelJsonSerialization(t *testing.T) {
	user := createUser()
	serialized, err := json.Marshal(user)
	assert.Nil(t, err)
	assert.Equal(t, `{"id":1,"name":"user_name","permissions":[1,2,3,4]}`, string(serialized))
}

func TestTableName(t *testing.T) {
	assert.Equal(t, "users", User{}.TableName())
}

func TestAfterFind(t *testing.T) {
	user := createUser()
	user.Permissions = nil
	user.AfterFind(nil)
	assert.Equal(t, []int{1, 2, 3, 4}, user.Permissions)
}

func TestBeforeSave(t *testing.T) {
	user := createUser()
	user.PermissionsDbField = ""
	user.BeforeSave(nil)
	assert.Equal(t, "1,2,3,4", user.PermissionsDbField)
}

func createUser() *User {
	return &User{
		ID:                 1,
		Name:               "user_name",
		Password:           "hash-for-password",
		RepeatedPassword:   "hash-for-repeated-password",
		Permissions:        []int{1, 2, 3, 4},
		PermissionsDbField: "1,2,3,4",
	}
}
