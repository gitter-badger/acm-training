package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"strconv"
)

// HasPermission - Checks whether the user has permission according to the
// permission list extracted from http headers.
func HasPermission(c *gin.Context, permission int) bool {
	if permissions, _ := c.Request.Header["Permission"]; len(permissions) > 0 {
		for _, item := range permissions {
			parsed, err := strconv.Atoi(item)
			if err == nil && parsed == permission {
				return true
			}
		}
	}
	return false
}

// CheckArgument - Checks the condition matches, otherwise the program will
// close according to the error. On the other hand, the cusotmized
// error message will be shown before exiting the program.
func CheckArgument(condition bool, messages ...interface{}) {
	if !condition {
		log.Fatal(messages...)
		os.Exit(1)
	}
}

// CheckIOError - Checks the IO error thrown by IO library reading or writing.
// This function will exit the program when EOF was read from the input,
// or any unexcepted IO exception thrown.
func CheckIOError(err error) {
	if err != nil {
		if err == io.EOF {
			fmt.Println("Interrupted, canceled and existing the tool.")
			os.Exit(1)
		} else {
			log.Fatal(err)
		}
	}
}
