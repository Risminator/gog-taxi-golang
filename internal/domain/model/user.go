package model

import "errors"

type User struct {
	UserId int
	Role   UserRole
}

type UserRole string

const (
	UnknownRole  UserRole = ""
	CustomerRole UserRole = "customer"
	DriverRole   UserRole = "driver"
)

func NewUser(id int, role UserRole) User {
	return User{id, role}
}

func UserRoleFromString(s string) (UserRole, error) {
	switch s {
	case "customer":
		return CustomerRole, nil
	case "driver":
		return DriverRole, nil
	}

	return UnknownRole, errors.New("unknown UserRole: " + s)
}
