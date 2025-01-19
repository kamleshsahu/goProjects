package main

import "fmt"

func getMessageStatus(timestamps []int, messages []string, k int) []bool {
	result := make([]bool, len(timestamps))
	queue := make([]int, 0, k)     // Queue to store the last k messages
	logMap := make(map[string]int) // Set to check duplicates quickly

	for i := 0; i < len(messages); i++ {
		message := messages[i]
		timestamp := timestamps[i]

		// Maintain the size of the queue and logMap to k
		for len(queue) > 0 && timestamps[queue[0]] < (timestamp-k) {
			oldest := queue[0]
			queue = queue[1:]                // Remove the oldest element from the queue
			delete(logMap, messages[oldest]) // Remove the oldest element from the logMap
		}

		prevTimestamp := logMap[message]
		if prevTimestamp == 0 {
			result[i] = true
		} else {
			result[i] = false
		}
		queue = append(queue, i)
		logMap[message] = timestamp
	}

	return result
}

func getMessageStatus2(timestamps []int, messages []string, k int) []bool {
	logMap := make(map[string]int)
	ans := make([]bool, len(timestamps))
	for i := 0; i < len(timestamps); i++ {
		timestamp := timestamps[i]
		message := messages[i]
		prevTimestamp := logMap[message]

		if prevTimestamp == 0 || prevTimestamp < timestamp-k {
			ans[i] = true
		} else {
			ans[i] = false
		}

		logMap[message] = timestamp
	}

	return ans
}

func main() {
	test1()
	test2()
}
func test1() {
	timestamps := []int{1, 4, 5, 10, 11, 14}
	messages := []string{"hello", "bye", "bye", "hello", "bye", "hello"}

	fmt.Println(getMessageStatus(timestamps, messages, 5))
}

func test2() {
	timestamps := []int{1, 1, 1, 11}
	messages := []string{"2", "2", "3", "2"}

	fmt.Println(getMessageStatus(timestamps, messages, 5))
}
