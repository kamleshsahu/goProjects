package main

import "fmt"

func maxAmount(initialCurrency string, pairs1 [][]string, rates1 []float64, pairs2 [][]string, rates2 []float64) float64 {

	graph := [2]map[string]map[string]float64{}
	visited = [2]map[string]map[string]bool{}
	graph[0] = make(map[string]map[string]float64)
	graph[1] = make(map[string]map[string]float64)
	visited[0] = make(map[string]map[string]bool)
	visited[1] = make(map[string]map[string]bool)

	for i, pair := range pairs1 {
		in, out := pair[0], pair[1]
		rate := rates1[i]
		if graph[0][in] == nil {
			graph[0][in] = make(map[string]float64)
			visited[0][in] = make(map[string]bool)
		}
		if graph[0][out] == nil {
			graph[0][out] = make(map[string]float64)
			visited[0][out] = make(map[string]bool)
		}
		graph[0][in][out] = float64(rate)
		graph[0][out][in] = 1.0 / float64(rate)
	}

	for i, pair := range pairs2 {
		in, out := pair[0], pair[1]
		rate := rates2[i]
		if graph[1][in] == nil {
			graph[1][in] = make(map[string]float64)
			visited[1][in] = make(map[string]bool)
		}
		if graph[1][out] == nil {
			graph[1][out] = make(map[string]float64)
			visited[1][out] = make(map[string]bool)
		}
		graph[1][in][out] = float64(rate)
		//graph[1][out][in] = 1.0 / float64(rate)
	}
	ans := 1.0
	fmt.Println(graph[0])
	fmt.Println(graph[1])

	ans = dfs(initialCurrency, graph, 0, initialCurrency, 1.0)
	return ans
}

var visited [2]map[string]map[string]bool

func dfs(initialCurr string, graph [2]map[string]map[string]float64, day int, curr string, value float64) float64 {
	ans := 0.0
	if initialCurr == curr {
		ans = max(ans, value)
	}
	fmt.Println(initialCurr, day, curr, value)
	if graph[day][curr] != nil {
		for targetCurr, rate := range graph[day][curr] {
			if targetCurr == initialCurr {
				ans = max(ans, value*rate)
				fmt.Println("reached target ", initialCurr, day, curr, value*rate)
				continue
			}
			if visited[day][curr][targetCurr] {
				fmt.Println("already visited ", day, curr, " to ", targetCurr, value*rate)
				continue;
			}
			visited[day][curr][targetCurr] = true
			fmt.Println("visiting  ", day, curr, " to ", targetCurr, value*rate)
			o1 := dfs(initialCurr, graph, day, targetCurr, value*rate)
			//visited[day][curr][targetCurr] = false
			ans = max(ans, o1)
		}
	}
	if day == 0 {
		o2 := dfs(initialCurr, graph, 1, curr, value)
		ans = max(ans, o2)
	}
	return ans
}

func main() {
	ans := maxAmount("EUR", [][]string{{"EUR", "USD"}, {"USD", "JPY"}}, []float64{2, 3}, [][]string{{"JPY", "USD"}, {"USD", "CHF"}, {"CHF", "EUR"}}, []float64{4, 5, 6})
	fmt.Println(ans)
}
