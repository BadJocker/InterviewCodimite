package main

import (
	"fmt"
	"strings"
)

func main() {

	var name string
	fmt.Println("Insert a name - ")
	fmt.Scan(&name)
	unique := hasAllUniqueCharacters(name)
	fmt.Println(unique)

}

func hasAllUniqueCharacters(name string) bool {
	fmt.Println("Name - ", name)
	var unique bool
	lowerCaseLetters := strings.ToLower(name)
	for _, lettersFirst := range lowerCaseLetters {
		for _, letterSecond := range lowerCaseLetters {
			if lettersFirst != letterSecond {
				unique = true
				continue
			} else {
				unique = false
				break
			}
		}
	}
	return unique
}
