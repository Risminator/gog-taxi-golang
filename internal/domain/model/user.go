package model

import "errors"

type User struct {
	UserId int
	Role   UserRole
}

type UserRole struct {
	slug string
}

func (r UserRole) String() string {
	return r.slug
}

var (
	UnknownRole  = UserRole{""}
	CustomerRole = UserRole{"customer"}
	DriverRole   = UserRole{"driver"}
)

func UserRoleFromString(s string) (UserRole, error) {
	switch s {
	case CustomerRole.slug:
		return CustomerRole, nil
	case DriverRole.slug:
		return DriverRole, nil
	}

	return UnknownRole, errors.New("unknown UserRole: " + s)
}
