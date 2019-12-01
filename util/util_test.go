package util

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHasPermission_HappyPath(t *testing.T) {
	// func HasPermission(c *gin.Context, permission int) bool
	result := HasPermission(&gin.Context{
		Request: &http.Request{
			Header: map[string][]string{
				"Permission": []string{"0"},
			},
		},
	}, 0)
	assert.True(t, result)
}

func TestHasPermission_NoPermissionFoudn(t *testing.T) {
	// func HasPermission(c *gin.Context, permission int) bool
	result := HasPermission(&gin.Context{
		Request: &http.Request{
			Header: map[string][]string{
				"Permission": []string{"0", "1"},
			},
		},
	}, 2)
	assert.False(t, result)
}

func TestCheckArgument_ConditionMeet(t *testing.T) {
	CheckArgument(true, "The exception will not be thrown.")
}

func TestCheckIOError_NoErrorOccurs(t *testing.T) {
	CheckIOError(nil)
}
