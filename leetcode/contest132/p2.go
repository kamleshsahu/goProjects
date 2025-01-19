package main

import (
	"sort"
)

//func main() {
//	skills := []int{1, 4, 3, 2, 5}
//	k := 3
//
//	fmt.Println(findWinningPlayer(skills, k))
//}

func findWinningPlayer(skills []int, k int) int {

	n := len(skills)

	players := make([]*Player, 0)

	for i, skill := range skills {
		players = append(players, &Player{Skill: skill, Idx: i, Wins: 0})
	}

	if k >= n {
		sort.Slice(players, func(i, j int) bool {
			return players[i].Skill < players[j].Skill
		})
		return players[n-1].Idx
	}

	i := 0

	for i < 2*n {

		next := (i + 1) % n
		if players[i].Skill > players[next].Skill {
			players[i], players[next] = players[next], players[i]
		}
		players[next].Wins++

		if players[next].Wins >= k {
			return players[next].Idx
		}

		i = next
	}

	return 0

}

type Player struct {
	Skill int
	Idx   int
	Wins  int
}
