package main

// User is a struct got by a repository
type User struct {
	ID    int
	Name  string
	Email string
}

// Repository is a sample interface
type Repository interface {
	Get(id int) *User
}

// UserRepository is a sample struct inherits Repository interface
type UserRepository struct {
}

// Get returns a user
func (r UserRepository) Get(id int) *User {
	for _, u := range DummyUsers {
		if u.ID == id {
			return u
		}
	}
	return nil
}

// DummyUsers is just a sample data
var DummyUsers = []*User{
	{
		ID:    1,
		Name:  "foo",
		Email: "foo@example.com",
	},
	{
		ID:    2,
		Name:  "bar",
		Email: "bar@example.com",
	},
}
