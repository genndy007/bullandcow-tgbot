package main

import (
	"math/rand"
	"time"
)

// Creates a number for player to guess
func GenerateNumber() int {
	// Very clever random number implementation :)
	result := 0
	rand.Seed(time.Now().UnixNano())
	allNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	plier := 1000
	for time := 0; time < 4; time++ {
		border := len(allNumbers)
		if time == 0 {
			border--
		}
		i := rand.Intn(border)
		result += allNumbers[i] * plier
		allNumbers = RemoveFromSlice(allNumbers, i)
		plier /= 10
	}
	return result

}

// Remove element from slice
func RemoveFromSlice(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

// Check if element is in slice
func IsInSlice(slice []rune, val rune) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// Method checks how many bulls and cows player has
func CheckBullsAndCows(guess, number string) [2]int {
	bullsAndCows := [2]int{0, 0}
	for i := 0; i < 4; i++ {
		if guess[i] == number[i] {
			bullsAndCows[0]++
		} else if IsInSlice([]rune(number), rune(guess[i])) {
			bullsAndCows[1]++
		}
	}

	return bullsAndCows
}
