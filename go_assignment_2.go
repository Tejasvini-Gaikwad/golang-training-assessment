package main

import (
	"fmt"
	"strings"
)

func getEntryKeys(entries map[string]bool) (keys []string) {
	for k, _ := range entries {
		keys = append(keys, k)
	}
	return
}

func main() {
	word := "product"

	// lookup for entries made by the user.
	entries := map[string]bool{}

	// list of "_" corrosponding to the number of letters in the word. [ _ _ _ _ _ ]
	placeholder := []string{}
	i := 0
	for i = 0; i < len(word); i++ {
		placeholder = append(placeholder, "_")
	}
	chances := len(word)
	for {
		// evaluate a loss! If user guesses a wrong letter or the wrong word, they lose a chance.
		userInput := strings.Join(placeholder, "")
		if chances == 0 && userInput != word {
			fmt.Println("Game over! please try again!")
			break
		}
		// evaluate a win!
		if userInput == word {
			fmt.Println("You win")
			break
		}
		// Console display
		fmt.Println(placeholder)                          // render the placeholder
		fmt.Printf("Chances: %d\n", chances)              // render the chances left
		fmt.Printf("Gueses %v\n ", getEntryKeys(entries)) // show the letters or words guessed till now.
		fmt.Print("Guess a letter or the word: ")

		// take the input
		str := ""
		fmt.Scanln(&str)
		if str == "" {
			fmt.Println("Please enter input")
			continue
		}
		if len(str) > 1 {
			if str == word {
				fmt.Println("You win!")
				break
			} else {
				entries[str] = true
				chances -= 1
				continue
			}

		}

		// compare and update entries, placeholder and chances.
		_, ok := entries[str]
		if ok {
			continue
		}
		entries[str] = true
		found := false
		splitted_word := strings.Split(word, "")
		for i, v := range splitted_word {
			if v == str {
				placeholder[i] = str
				found = true

			}
		}
		if !found {
			chances -= 1

		}
	}
}
