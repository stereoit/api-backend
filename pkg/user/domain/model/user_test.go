package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewUser(t *testing.T) {
	assert := assert.New(t)
	testID := "1"
	testEmail := "test@email.cz"
	testFirstName := "Václav"
	testLastName := "Havel"

	user := NewUser(testID, testEmail)
	assert.NotNil(user, "user should not be nil")

	got := user.GetID()
	assert.Equal(testID, got, "user ID should match")

	got = user.GetEmail()
	assert.Equal(testEmail, got, "user email should match")

	got = user.GetFirstName()
	assert.Equal("", got, "first name should be empty")

	user.SetFirstName(testFirstName)
	got = user.GetFirstName()
	assert.Equal(testFirstName, got, "first name should match")

	got = user.GetLastName()
	assert.Equal("", got, "last name should be empty")

	user.SetLastName(testLastName)
	got = user.GetFirstName()
	assert.Equal(testLastName, got, "last name should match")

}
