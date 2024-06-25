package main

import (
	"encoding/json" 
	"fmt"
	"flag"
	"time"
	"strings"
	"net/http"
	"io/ioutil"
)

var dev = flag.Bool("dev", false, "dev server")

func getEntryKeys(entries map[string]bool) (keys []string) {
	for k, _ := range entries {
		keys = append(keys, k)
	}
	return
}

func getWord() string {
	if *dev {
		return "product";
	}
	resp, err := http.Get("https://random-word-api.herokuapp.com/word?number=10")
	
	if err != nil {
		return "product";
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", body)
	words := []string{};
	err = json.Unmarshal(body, &words)
	if err != nil {
		return "product";
	}
	for i,v := range words{
		if len(v)>4 && len(v) <8 {
			return words[i]
		} 
	}
	
	return words[0]
	
}

func main() {
	flag.Parse()
	word := getWord()

	// lookup for entries made by the user.
	entries := map[string]bool{}

	// list of "_" corrosponding to the number of letters in the word. [ _ _ _ _ _ ]
	placeholder := []string{}
	i := 0
	for i = 0; i < len(word); i++ {
		placeholder = append(placeholder, "_")
	}
	
	t := time.NewTimer(5 * time.Second);
	result := make(chan bool)
	//expired := <- t.C
	//fmt.Println(expired, "ss")
	chances := len(word)
	go func () {
	for {
		// evaluate a loss! If user guesses a wrong letter or the wrong word, they lose a chance.
		userInput := strings.Join(placeholder, "")
		if chances == 0 && userInput != word {
			result <- false
			fmt.Println("Game over! please try again!")
			return
		}
		// evaluate a win!
		if userInput == word {
			result <- true
			fmt.Println("You win")
			return
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
				result <- true
				fmt.Println("You win!")
				return
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
}()
	//for {
		select {
			case <-result:
			fmt.Println("...")
			break;
			//goto END
			case <-t.C:
			fmt.Println("TimedOut...")
			break;
			//goto END
		}
	//}
//END:
	fmt.Println("Play again!")
}
