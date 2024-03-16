package repository

import "context"

type IUserRepository interface {
	Save(context.Context, *User) (*User, error)
	FindByPhoneNumber(context.Context, string) (*User, bool, error)
	FindByID(context.Context, int) (*User, bool, error)
	Update(context.Context, *User) (*User, error)
}
