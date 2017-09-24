package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Introduce un texto: ")
	input,_ := reader.ReadString('\n')

	CountVowels(input)
}

func CountVowels(input string) {
	input = strings.ToUpper(input)
	var sliceInput = strings.Split(input, "")

	var data = map[string]int{
		"A": 0,
		"E": 0,
		"I": 0,
		"O": 0,
		"U": 0,
	}

	for _, vowel := range sliceInput {
		if _, ok := data[vowel]; ok {
			data[vowel]++
		}
	}

	fmt.Println(BuildResult(data))
}

func BuildResult(data map[string]int)  string{
	var output string
	for vowel, count := range data {
		var plural ="vez"
		if count > 1 || count == 0 {
			plural = "veces"
		}

		output += fmt.Sprintf("* %s esta %d %s\n", vowel, count, plural)
	}

	return output
}