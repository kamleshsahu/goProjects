package service

type IBalanceSheetService interface {
	AddToSheet(paidBy int, owes int, amount int)
	ReverseFromSheet(paidBy int, owes int, amount int)
	GetBalanceSheet(uid int) map[int]int
}
