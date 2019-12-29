package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	got = user.GetLastName()
	assert.Equal(testLastName, got, "last name should match")

}

func Test_Validate(t *testing.T) {
	assert := assert.New(t)
	//
	testUser := &User{}
	err := testUser.Validate()
	assert.NotNil(err, "id and email cannot be blank")

	testUser = &User{
		id:    "8e492e66-9af8-48ab-a22d-61cbaaf333fa",
		email: "bad-email",
	}
	err = testUser.Validate()
	assert.NotNil(err, "email must be valid address")

	testUser = &User{
		id:    "8e492e66-9af8-48ab-a22d-61cbaaf333fa",
		email: "user@example.com",
	}
	err = testUser.Validate()
	assert.Nil(err, "correct user should validate")
}
