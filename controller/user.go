package user

import (
	"github.com/gin-gonic/gin"
	"github.com/uestc-acm/acm-training/model"
	"net/http"
)

// AddUser is a API to add a user into the system.
// This API only allows already signed in administrators to create users.
// TODO(ruinshe): add permission check.
func AddUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Permissions == nil {
		user.Permissions = make([]uint32, 0)
	}

	user.ID = 1
	// TODO(ruinshe): Save user here.
	c.JSON(http.StatusCreated, user)
}

// GetUsers is a API to get all users from the database.
// TODO(ruinshe): consider add season concept, or other group concept.
// TODO(ruinshe): add pagination support.
func GetUsers(c *gin.Context) {
	// TODO(ruinshe): Get users from the storage.
	users := make([]model.User, 0)
	c.JSON(http.StatusOK, gin.H{"users": users})
}
