package impl

import (
	"awesomeProject/lld/splitwise/model"
	"awesomeProject/lld/splitwise/service"
)

type bss struct {
	balanceSheet model.UserToUserBalance
}

func (b *bss) AddToSheet(paidBy int, owes int, amount int) {
	if paidBy == owes {
		return
	}
	if _, ok := b.balanceSheet[paidBy]; !ok {
		b.balanceSheet[paidBy] = make(map[int]int)
	}
	if _, ok := b.balanceSheet[owes]; !ok {
		b.balanceSheet[owes] = make(map[int]int)
	}
	b.balanceSheet[paidBy][owes] += amount
	b.balanceSheet[owes][paidBy] -= amount
}

func (b *bss) ReverseFromSheet(paidBy int, owes int, amount int) {
	if paidBy == owes {
		return
	}
	if _, ok := b.balanceSheet[paidBy]; !ok {
		b.balanceSheet[paidBy] = make(map[int]int)
	}
	if _, ok := b.balanceSheet[owes]; !ok {
		b.balanceSheet[owes] = make(map[int]int)
	}

	b.balanceSheet[paidBy][owes] -= amount
	b.balanceSheet[owes][paidBy] += amount
}

func (b *bss) GetBalanceSheet(uid int) map[int]int {
	return b.balanceSheet[uid]
}

func NewBalanceSheet() service.IBalanceSheetService {
	return &bss{balanceSheet: make(model.UserToUserBalance)}
}
