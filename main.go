package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var inputReader = bufio.NewReader(os.Stdin)

var dictionary = []string{
	"Golang",
	"Programming",
	"Language",
	"Turkiye",
	"Test",
}

// Implement the command 'hint' - show to the user a random unguessed letter. limit the user to one hint only

func main() {
	rand.Seed(time.Now().UnixNano())

	targetWord := getRandomWord()
	guessedLetters := initializeGuessedWords(targetWord)
	hangmanState := 0

	for !isGameOver(targetWord, guessedLetters, hangmanState) {
		printGameState(targetWord, guessedLetters, hangmanState)
		input := readInput()
		if len(input) != 1 {
			fmt.Println("\n\nInvalid input. Please use letters only!!!")
			continue
		}

		letter := rune(input[0])
		if isCorrectGuess(targetWord, letter) {
			if guessedLetters[letter] == true {
				fmt.Println("You've already used that letter")
			}
			guessedLetters[letter] = true
		} else {
			hangmanState++
		}
	}
	printGameState(targetWord, guessedLetters, hangmanState)
	fmt.Println("!!! GAME OVER !!!")
	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Println("YOU WIN")
	} else if isHangmanCompleted(hangmanState) {
		fmt.Println("YOU LOSE")
	} else {
		panic("Invalid state. Game is over and there is no winner!")
	}
}

func getRandomWord() string {
	return dictionary[rand.Intn(len(dictionary))]
}

func initializeGuessedWords(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	return guessedLetters
}

func isGameOver(targetWord string, guessedLetters map[rune]bool, hangmanstate int) bool {
	return isWordGuessed(targetWord, guessedLetters) || isHangmanCompleted(hangmanstate)
}

func isWordGuessed(targetWords string, guessedLetters map[rune]bool) bool {
	for _, ch := range targetWords {
		if !guessedLetters[unicode.ToLower(ch)] {
			return false
		}
	}
	return true
}

func isHangmanCompleted(hangmanState int) bool {
	return hangmanState >= 9
}

func printGameState(targetWord string, guessedLetters map[rune]bool, hangmanState int) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println()
	fmt.Println(getHangmanDrawing(hangmanState))
}

func getWordGuessingProgress(targetWord string, guessedLetters map[rune]bool) string {
	result := ""
	for _, ch := range targetWord {
		if ch == ' ' {
			result += ""
		} else if guessedLetters[unicode.ToLower(ch)] == true {
			result += fmt.Sprintf("%c", ch)
		} else {
			result += "_"
		}
		result += " "
	}
	return result
}

func getHangmanDrawing(hangmanState int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("states/hangman%d", hangmanState))

	if err != nil {
		panic(err)
	}
	return string(data)
}

func readInput() string {
	fmt.Print("> ")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(input)
}

func isCorrectGuess(targetWord string, letter rune) bool {
	return strings.ContainsRune(targetWord, letter)
}
