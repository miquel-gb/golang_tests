package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	words := strings.Fields(s)
	for _, w := range words {
		wordMap[w]++
	}
	return wordMap
}

func fibonacci() func() int {
	num1, num2 := 0, 1
	return func() int {
		num1, num2 = num2, num1+num2
		return num1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	/*fmt.Println(WordCount("I am I Go! am learning Go!"))*/
	/*m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])*/
}
