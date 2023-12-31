package dices

import (
	"math/rand"
	"time"
)

// rollDice roll's dice of given number of walls (accepts D4, D6, D8, D12, D20)
func rollDice(number int) int {
	randomSource := rand.NewSource(time.Now().UnixNano())

	randomGenerator := rand.New(randomSource)

	randomNumber := randomGenerator.Intn(number) + 1

	return randomNumber
}

// RollD4 roll's D4 dice
func RollD4() int {
	return rollDice(4)
}

// RollD6 roll's D6 dice
func RollD6() int {
	return rollDice(6)
}

// RollD8 roll's D8 dice
func RollD8() int {
	return rollDice(8)
}

// RollD8 roll's D10 dice
func RollD10() int {
	return rollDice(10)
}

// RollD12 roll's D12 dice
func RollD12() int {
	return rollDice(12)
}

// RollD20 roll's D20 dice
func RollD20() int {
	return rollDice(20)
}

// RollD100 roll's D100 dice
func RollD100() int {
	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(randomSource)

	poolTens := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}

	randomIndex := randomGenerator.Intn(len(poolTens))
	randomNumber := poolTens[randomIndex]

	randomOnesNumber := randomGenerator.Intn(10) + 1

	return randomNumber + randomOnesNumber
}
