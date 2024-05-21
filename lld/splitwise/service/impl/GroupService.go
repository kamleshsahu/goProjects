package impl

import (
	"awesomeProject/lld/splitwise/model"
	service2 "awesomeProject/lld/splitwise/service"
)

type groupService struct {
	bss service2.IBalanceSheetService
	ss  service2.ISplitService
	gm  service2.IGMService
}

func (g *groupService) getGroup(gid int) *model.Group {
	return g.gm.GetGroup(gid)
}

func (g *groupService) AddExpense(gid int, expense model.Expense) {
	if expense.SplitType == model.EQUAL {
		userMap := make(map[int]int)
		for userId := range g.gm.GetGroup(gid).Users {
			userMap[userId] = 0
		}

		expense.Splits = g.ss.SplitAmount(expense.Total, expense.SplitType, userMap)
	}

	grp := g.getGroup(gid)
	grp.Expenses = append(grp.Expenses, expense)
	g.addToBalanceSheet(expense)
}

func (g *groupService) addToBalanceSheet(expense model.Expense) {
	for _, split := range expense.Splits {
		g.bss.AddToSheet(expense.PaidBy, split.UserId, split.Amount)
	}
}

func (g *groupService) reverseFromBalanceSheet(expense model.Expense) {
	for _, split := range expense.Splits {
		g.bss.ReverseFromSheet(expense.PaidBy, split.UserId, split.Amount)
	}
}

func (g *groupService) AddUser(gid, uid int) {
	grp := g.getGroup(gid)
	grp.Users[uid] = true
}

func (g *groupService) RemoveUser(gid, uid int) {
	grp := g.getGroup(gid)
	users := grp.Users
	delete(users, uid)
}

func (g *groupService) DeleteExpense(gid, eid int) {
	expense := g.getGroup(gid).Expenses[eid]
	expenses := g.getGroup(gid).Expenses
	expenses = append(expenses[:eid], expenses[eid+1:]...)
	g.reverseFromBalanceSheet(expense)
}

func (g *groupService) UpdateExpense(gid, eid int, expense model.Expense) {
	g.reverseFromBalanceSheet(g.getGroup(gid).Expenses[eid])
	g.getGroup(gid).Expenses[eid] = expense
	g.addToBalanceSheet(g.getGroup(gid).Expenses[eid])
}

func NewGroupService(bss service2.IBalanceSheetService, gm service2.IGMService) service2.IGroupService {
	return &groupService{bss: bss, ss: NewSplitService(), gm: gm}
}
