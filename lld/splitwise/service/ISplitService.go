package service

import (
	"awesomeProject/lld/splitwise/model"
)

type ISplitService interface {
	SplitAmount(total int, splittype model.SPLITTYPE, splits map[int]int) []model.Split
}
