package main

import (
	"fmt"
	"sort"
)

func main() {
	accounts := [][]string{{"David", "David0@m.co", "David1@m.co"}, {"David", "David2@m.co", "David3@m.co"}, {"David", "David3@m.co", "David1@m.co"}}

	//accounts := [][]string{{"David","David0@m.co","David1@m.co"},{"David","David3@m.co","David4@m.co"},{"David","David4@m.co","David5@m.co"},{"David","David2@m.co","David3@m.co"},{"David","David1@m.co","David2@m.co"}}
	accountsMerge(accounts)
}

func accountsMerge(accounts [][]string) [][]string {
	// emailToAccountList

	emailToAccountList := make(map[string][]int)
	parents := make([]int, len(accounts))
	for i, account := range accounts {
		for _, email := range account[1:] {
			if emailToAccountList[email] == nil {
				emailToAccountList[email] = make([]int, 0)
			}
			emailToAccountList[email] = append(emailToAccountList[email], i)
		}
		parents[i] = i
	}
	fmt.Println(parents)

	// parentMap
	fmt.Println(emailToAccountList)
	for _, accountIds := range emailToAccountList {
		for _, accountId := range accountIds[1:] {
			fmt.Println("union find :", accountIds[0], accountId)
			union(parents, accountIds[0], accountId)
		}
	}
	fmt.Println(parents)

	// accountToEmailMap
	_accounts := make(map[int]*Account, len(accounts))
	for i, parent := range parents {
		if _accounts[parent] == nil {
			_accounts[parent] = &Account{
				Id:   parent,
				Name: accounts[i][0],
			}
		}

		(*_accounts[parent]).Emails = append((*_accounts[parent]).Emails, accounts[i][1:]...)
	}

	ans := make([][]string, 0)

	for _, account := range _accounts {

		name := (*account).Name
		arr := []string{}
		arr = append(arr, name)
		arr = append(arr, unique((*account).Emails)...)
		ans = append(ans, arr)
	}

	return ans

}

func unique(emails []string) (ans []string) {
	emailSet := make(map[string]bool)

	for _, email := range emails {
		emailSet[email] = true
	}

	for key, _ := range emailSet {
		ans = append(ans, key)
	}

	sort.Strings(ans)
	return ans
}

type Account struct {
	Id     int
	Name   string
	Emails []string
}

func union(parents []int, i, j int) {
	p1, p2 := find(parents, i), find(parents, j)
	if p1 > p2 {
		p1, p2 = p2, p1
	}
	parents[p2] = p1
	parents[j] = p1
}

func find(parents []int, idx int) int {
	if parents[idx] == idx {
		return idx
	}
	parents[idx] = find(parents, parents[idx])

	return parents[idx]
}
