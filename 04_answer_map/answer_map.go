// 完全に自分で作りました＞＜
package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordcount := make(map[string]int)
	str := strings.Fields(s)
	fmt.Printf("Type=%T , Value = %s\n", str, str)
	for _, i := range str {
		wordcount[i] = wordcount[i] + 1
		fmt.Println(i, wordcount[i])
	}

	return wordcount
	//	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
