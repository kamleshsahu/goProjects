package service

import (
	"awesomeProject/lld/splitwise/model"
)

type IUserService interface {
	RegisterUser(user model.User) int
	DeleteUser(uid int)
	AddToGroup(uid, gid int)
}
