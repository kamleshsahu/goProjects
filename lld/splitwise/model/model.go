package model

import "time"

const (
	CUSTOM = iota
	EQUAL
)

type SPLITTYPE int

type Split struct {
	Amount int
	UserId int
}

type Expense struct {
	Eid       int
	Splits    []Split
	PaidBy    int
	TimeStamp time.Time
	Total     int
	IsDeleted bool
	SplitType SPLITTYPE
}

type User struct {
	Uid    int
	Name   string
	Groups map[int]bool
}

type Group struct {
	Expenses     []Expense
	Users        map[int]bool
	TotalExpense int
	IsDeleted    bool
}

type UserToUserBalance map[int]map[int]int
