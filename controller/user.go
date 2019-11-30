package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/uestc-acm/acm-training/db"
	"github.com/uestc-acm/acm-training/model"
	"github.com/uestc-acm/acm-training/util"
	"net/http"
)

// AddUser is a API to add a user into the system.
// This API only allows already signed in administrators to create users.
// TODO(ruinshe): use JWT for permission check.
func AddUser(c *gin.Context) {
	if !util.HasPermission(c, model.PermissionAddUser) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "User has no permission to create new user.",
		})
		return
	}
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(user.Name) > 32 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("The field length `name` is too long: %v", len(user.Name)),
		})
		return
	}

	// TODO(ruinshe): check exsistence.
	db.DB().Create(&user)
	c.JSON(http.StatusCreated, user)
}

// GetUsers is a API to get all users from the database.
// TODO(ruinshe): consider add season concept, or other group concept.
// TODO(ruinshe): add pagination support.
func GetUsers(c *gin.Context) {
	// TODO(ruinshe): Get users from the storage.
	var users []model.User
	db.DB().Limit(50).Offset(0).Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}
