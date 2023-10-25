// models/user.go
package models

type User struct {
	ID    int
	Name  string
	Email string
}

// For the sake of this example, we'll return some mock data
func GetUsers() []User {
	return []User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
	}
}
