package service

import (
	"awesomeProject/lld/splitwise/model"
)

type IGroupService interface {
	AddExpense(gid int, expense model.Expense)
	AddUser(gid, uid int)
	RemoveUser(gid, uid int)
	DeleteExpense(gid, eid int)
	UpdateExpense(gid, eid int, expense model.Expense)
}
