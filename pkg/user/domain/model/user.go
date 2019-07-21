package model

// User represents single user
type User struct {
	id        string
	email     string
	firstName string
	lastName  string
}

// NewUser returns new instance of the user
func NewUser(id, email, firstName, lastName string) *User {
	return &User{
		id:        id,
		email:     email,
		firstName: firstName,
		lastName:  lastName,
	}
}

// GetID returns ID of the user
func (u *User) GetID() string {
	return u.id
}

// GetEmail returns the user's email
func (u *User) GetEmail() string {
	return u.email
}

// GetFirstName returns the first name of the user
func (u *User) GetFirstName() string {
	return u.firstName
}

// GetLastName returns the last name of the user
func (u *User) GetLastName() string {
	return u.lastName
}
