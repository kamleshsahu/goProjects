package main

import (
	"strconv"
)

//func main() {
//	fmt.Println(clearDigits("a1ab23ffgagag232"))
//}

func clearDigits(s string) string {

	arr := []byte(s)

	ans := make([]byte, 0)

	for _, val := range arr {

		_, err := strconv.Atoi(string(val))

		if err != nil {
			ans = append(ans, byte(val))
		} else if len(ans) > 0 {
			ans = ans[:len(ans)-1]
		}

	}

	return string(ans)
}
