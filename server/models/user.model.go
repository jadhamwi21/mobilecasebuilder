package models

type RoleType string

const (
	Client RoleType = "client"
	Staff  RoleType = "staff"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	Role      RoleType
}
