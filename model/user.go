package model

import (
	"github.com/jinzhu/gorm"
	"github.com/uestc-acm/acm-training/db"
)

const (
	// PermissionAddUser describes a permission that the user is allowed
	// to create a new non administrator user.
	PermissionAddUser int = 0
)

// User is the basic user definition of the system, including authorization
// informaiton and role owned by the user.
type User struct {
	// The unique identifier of the entity, for DB storing and indexing.
	ID uint64 `json:"id"`
	// The name of the user, in this case, we use real name for internal
	// management purpose. For general usage, it maybe a nickname or any
	// fake string.
	Name string `json:"name" binding:"required" gorm:"type:varchar(32);unique_index"`
	// The encrypted password of the user by cryptographic library.
	Password         string `json:"-"`
	RepeatedPassword string `json:"-" gorm:"-"`
	// The permissions of the user, for API access control.
	Permissions []int `json:"permissions" sql:"-"`
	// the database []int field for persistence.
	PermissionsDbField string `json:"-" gorm:"column:permissions"`
}

// TableName - The name of the table for users.
func (User) TableName() string {
	return "users"
}

// AfterFind - The action to be performed after find.
func (user *User) AfterFind(scope *gorm.Scope) (err error) {
	user.Permissions, err = db.BindIntsAfterFind(scope, user.PermissionsDbField)
	return
}

// BeforeSave - The action to be performed before save.
func (user *User) BeforeSave(scope *gorm.Scope) (err error) {
	user.PermissionsDbField, err = db.BindIntsBeforeSave(scope, user.Permissions)
	return
}
