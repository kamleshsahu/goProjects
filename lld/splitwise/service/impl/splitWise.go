package impl

import (
	"awesomeProject/lld/splitwise/model"
	"encoding/json"
	"fmt"
	"time"
)

func Splitwise() {

	NewSplitService()

	//splits := ss.SplitAmount(100, model.EQUAL, nil)

	us := NewUserService()

	kamlesh := us.RegisterUser(model.User{
		Name:   "kamlesh",
		Groups: make(map[int]bool),
	})

	bss := NewBalanceSheet()
	gm := NewGMService(us)

	gs := NewGroupService(bss, gm)

	tikesh := us.RegisterUser(model.User{
		Name:   "tikesh",
		Groups: make(map[int]bool),
	})

	nilesh := us.RegisterUser(model.User{
		Name:   "nilesh",
		Groups: make(map[int]bool),
	})
	gid := gm.CreateGroup([]int{kamlesh, tikesh, nilesh})

	expense1 := model.Expense{
		PaidBy:    kamlesh,
		TimeStamp: time.Time{},
		Total:     100,
		IsDeleted: false,
		SplitType: model.EQUAL,
	}

	expense2 := model.Expense{
		PaidBy:    tikesh,
		TimeStamp: time.Time{},
		Total:     100,
		IsDeleted: false,
		SplitType: model.EQUAL,
	}

	expense3 := model.Expense{
		PaidBy:    nilesh,
		TimeStamp: time.Time{},
		Total:     100,
		IsDeleted: false,
		SplitType: model.EQUAL,
	}

	gs.AddExpense(gid, expense1)
	gs.AddExpense(gid, expense2)

	gs.RemoveUser(gid, kamlesh)
	gs.AddExpense(gid, expense3)
	tbss := bss.GetBalanceSheet(tikesh)

	fmt.Println("tikesh", tbss)

	kbss := bss.GetBalanceSheet(kamlesh)
	fmt.Println("kamlesh", kbss)

	nbss := bss.GetBalanceSheet(nilesh)
	fmt.Println("nilesh", nbss)

	jsonStr, _ := json.Marshal(gm.GetGroup(gid))
	fmt.Println(string(jsonStr))

}
