package main

import "fmt"

//You have 50 bitcoins to distribute to
// 10 users: Matthew, Sarah, Augustus, Heidi, Emilie, Peter, Giana, Adriano, Aaron, Elizabeth
// The coins will be distributed based on the vowels contained in each name where:

// a: 1 coin e: 1 coin i: 2 coins o: 3 coins u: 4 coins

// and a user can’t get more than 10 coins.
// Print a map with each user’s name and the amount of coins distributed.
// After distributing all the coins, you should have 2 coins left.

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	fmt.Println(distribution)
	coinsForUser := func(user string) int {
		expectedCoins := 0
		for i := 0; i < len(user); i++ {
			switch string(user[i]) {
			case "a", "A":
				expectedCoins++
			case "e", "E":
				expectedCoins++
			case "i", "I":
				expectedCoins = expectedCoins + 2
			case "o", "O":
				expectedCoins = expectedCoins + 3
			case "u", "U":
				expectedCoins = expectedCoins + 4
			}
		}
		return expectedCoins
	}

	for _, user := range users {
		if coins <= 0 {
			break
		}
		expectedCoins := coinsForUser(user)

		if expectedCoins > 10 {
			expectedCoins = 10
		}
		coins = coins - expectedCoins
		distribution[user] = expectedCoins
	}
	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
}
