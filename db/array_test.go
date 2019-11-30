package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBindIntsAfterFind_HappyPath(t *testing.T) {
	values, err := BindIntsAfterFind(nil, "1,2,3,4,5")
	assert.Equal(t, []int{1, 2, 3, 4, 5}, values)
	assert.Nil(t, err)
}

func TestBindIntsAfterFind_EmptyArray(t *testing.T) {
	values, err := BindIntsAfterFind(nil, "")
	assert.Equal(t, make([]int, 0), values)
	assert.Nil(t, err)
}

func TestBindIntsAfterFind_InvalidValue(t *testing.T) {
	values, err := BindIntsAfterFind(nil, "invalid")
	assert.Equal(t, make([]int, 0), values)
	assert.Equal(
		t,
		"strconv.Atoi: parsing \"invalid\": invalid syntax",
		err.Error())
}

func TestBindIntsBeforeSave_HappyPath(t *testing.T) {
	dbField, err := BindIntsBeforeSave(nil, []int{1, 2, 3, 4, 5})
	assert.Equal(t, "1,2,3,4,5", dbField)
	assert.Nil(t, err)
}

func TestBindIntsBeforeSave_EmptyArray(t *testing.T) {
	dbField, err := BindIntsBeforeSave(nil, make([]int, 0))
	assert.Equal(t, "", dbField)
	assert.Nil(t, err)
}

func TestBindIntsBeforeSave_NilArray(t *testing.T) {
	dbField, err := BindIntsBeforeSave(nil, nil)
	assert.Equal(t, "", dbField)
	assert.Nil(t, err)
}
