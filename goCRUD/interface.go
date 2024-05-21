package main

type LibraryServiceI interface {
	CreateUser(user User) (*User, CustomError)
	UpdateUser(id int, user User) (updatedUser *User, customError CustomError)
	DeleteUser(id int) CustomError
	GetUsers(interface{}) ([]User, CustomError)
}

type LibraryRepoI interface {
	CreateUser(user User) (*User, CustomError)
	UpdateUser(id int, user User) (updatedUser *User, customError CustomError)
	DeleteUser(id int) CustomError
	GetUsers(interface{}) ([]User, CustomError)
}
