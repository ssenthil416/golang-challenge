package dataservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * TODO: FIX TestMain so it call the test functions correctly
 */
func TestMain(t *testing.T) {
	checkInit(t, nil)
	checkGet(t, nil)
	checkDelete(t, nil)
}

/*
 * DO NOT MODIFY THIS TEST FUNCTION
 */
func checkInit(t *testing.T, data DumbDictionary) {
	assert.NotNil(t, data)
	err := data.Init("../data.gob")
	assert.Nil(t, err)
}

/*
 * DO NOT MODIFY THIS TEST FUNCTION
 */
func checkGet(t *testing.T, data DumbDictionary) {
	assert.NotNil(t, data)
	length, err := data.Get("Mike Tyson")
	assert.Nil(t, err)
	assert.Equal(t, int32(10), length)
}

/*
 * DO NOT MODIFY THIS TEST FUNCTION
 */
func checkDelete(t *testing.T, data DumbDictionary) {
	assert.NotNil(t, data)
	_, err := data.Delete("Mike Tyson")
	assert.Nil(t, err)
	_, err = data.Get("Mike Tyson")
	assert.NotNil(t, err)
}
