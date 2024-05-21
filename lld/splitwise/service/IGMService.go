package service

import (
	"awesomeProject/lld/splitwise/model"
)

type IGMService interface {
	CreateGroup(users []int) int
	GetGroup(gid int) *model.Group
}
