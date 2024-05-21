package impl

import (
	"awesomeProject/lld/splitwise/model"
	"awesomeProject/lld/splitwise/service"
)

type splitService struct {
}

func (s *splitService) SplitAmount(total int, splittype model.SPLITTYPE, userToSplit map[int]int) []model.Split {
	switch splittype {
	case model.CUSTOM:
		return s.customSplit(total, userToSplit)
	case model.EQUAL:
		return s.equalSplit(total, userToSplit)
	}
	return nil
}

func (s *splitService) equalSplit(total int, userToSplit map[int]int) (splitList []model.Split) {
	userCount := len(userToSplit)
	split := total / userCount
	for user, _ := range userToSplit {
		splitList = append(splitList, model.Split{Amount: split, UserId: user})
	}

	return splitList
}

func (s *splitService) customSplit(total int, userToSplit map[int]int) (splitList []model.Split) {

	actualTotal := 0
	for user, split := range userToSplit {
		actualTotal += split
		splitList = append(splitList, model.Split{Amount: split, UserId: user})
	}
	if total != actualTotal {
		return nil
	}
	return splitList
}

func NewSplitService() service.ISplitService {
	return &splitService{}
}
