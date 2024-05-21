package impl

import (
	"awesomeProject/lld/splitwise/model"
	service2 "awesomeProject/lld/splitwise/service"
)

type gmService struct {
	groupMap    map[int]*model.Group
	nextGroupId int
	userService service2.IUserService
}

func (g *gmService) GetGroup(gid int) *model.Group {
	return g.groupMap[gid]
}

func (g *gmService) CreateGroup(users []int) int {
	gid := g.nextGroupId
	userMap := make(map[int]bool)
	for _, userId := range users {
		userMap[userId] = true
		g.userService.AddToGroup(gid, userId)
	}
	g.groupMap[gid] = &model.Group{
		Expenses:     make([]model.Expense, 0),
		Users:        userMap,
		TotalExpense: 0,
		IsDeleted:    false,
	}

	g.nextGroupId++
	return gid
}

func NewGMService(us service2.IUserService) service2.IGMService {
	return &gmService{groupMap: make(map[int]*model.Group), userService: us}
}
