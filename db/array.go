package db

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

const arraySplitChar = ","

// BindIntsAfterFind - Binds the ints array (i.e. []int) to a string field
// for sqlite perissting.
func BindIntsAfterFind(_ *gorm.Scope, dbField string) (values []int, err error) {
	if dbField == "" {
		values = make([]int, 0)
		return
	}
	arr := strings.Split(dbField, arraySplitChar)
	values = make([]int, len(arr))
	for i := range values {
		values[i], err = strconv.Atoi(arr[i])
		if err != nil {
			values = make([]int, 0)
			return
		}
	}

	return
}

// BindIntsBeforeSave - Binds the ints array (i.e. []int) to a string field
// for sqlite data displaying.
func BindIntsBeforeSave(_ *gorm.Scope, values []int) (dbField string, err error) {
	if values == nil {
		dbField = ""
		return
	}
	var blocks []string = make([]string, len(values))
	for i := range values {
		blocks[i] = strconv.Itoa(values[i])
	}

	dbField = strings.Join(blocks, arraySplitChar)
	return
}
