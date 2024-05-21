package impl

import (
	"awesomeProject/lld/splitwise/model"
	"awesomeProject/lld/splitwise/service"
)

type userService struct {
	users      map[int]model.User
	nextUserId int
}

func (u *userService) RegisterUser(user model.User) int {

	uid := u.nextUserId
	u.users[uid] = user
	u.nextUserId++
	return uid
}

func (u *userService) DeleteUser(userId int) {
	delete(u.users, userId)
}

func (u *userService) AddToGroup(userId, gid int) {
	_, ok := u.users[userId]

	if !ok {
		return
	}

	u.users[userId].Groups[gid] = true
}

func NewUserService() service.IUserService {
	return &userService{
		users: make(map[int]model.User),
	}
}
