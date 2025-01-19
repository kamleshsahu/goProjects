package main

import (
	"awesomeProject/designPatterns/strategy/strategy"
	"fmt"
)

func main() {
	tp := NewTextProcessor(strategy.CSVStrategy())
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp)

	//tp.Reset()
	//tp.SetOutputFormat(Html)
	//tp.AppendList([]string{"foo", "bar", "baz"})
	//fmt.Println(tp)
}
