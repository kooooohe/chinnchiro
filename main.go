package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func rollDice() []int {
	dice := []int{rand.Intn(6) + 1, rand.Intn(6) + 1, rand.Intn(6) + 1}
	sort.Ints(dice)
	return dice
}

func getPayout(roll []int) (int, int) {
	// fmt.Println(roll)
	// 111
	if roll[0] == 1 && roll[1] == 1 && roll[2] == 1 {
		return 5, 0
	}
	// 456
	if roll[0] == 4 && roll[1] == 5 && roll[2] == 6 {
		return 2, 0
	}
	// 222,333,444,555,666
	if roll[0] == roll[1] && roll[1] == roll[2] {
		return 3, 0
	}

	// 123
	if roll[0] == 1 && roll[1] == 2 && roll[2] == 3 {
		return -2, 0
	}

	// 112,221,...
	if roll[0] == roll[1] {
		return 1, roll[2]
	}
	if roll[1] == roll[2] {
		return 1, roll[0]
	}
	if roll[0] == roll[2] {
		return 1, roll[1]
	}
	// 134,...
	return 0, 0
}

func decideWinner(parentRoll, childRoll []int) string {
	parentPayout, pDiceN := getPayout(parentRoll)
	childPayout, cDiceN := getPayout(childRoll)

	if parentPayout == 1 && childPayout == 1 {
		if pDiceN >= cDiceN {
			return "parent"
		}
	}

	if parentPayout >= childPayout {
		return "parent"
	}
	if parentPayout < childPayout {
		return "child"
	}

	return "parent"
}

func playGame() (float64, float64) {
	dMoeny := 100
	dLoopN := 1
	parentMoney := dMoeny
	childMoney := dMoeny

	for i := 0; i < dLoopN; i++ {
		pDices := []int{}
		for j := 0; j < 3; j++ {
			pDices = rollDice()
			//pDices = []int{2,2,3}
			//pDices = []int{1,2,3}
			pR, _ := getPayout(pDices)
			fmt.Print("pR: ")
			fmt.Println(pR)
			fmt.Println(pDices)
			if pR != 0 {
				break
			}
		}

		cDices := []int{}
		for j := 0; j < 3; j++ {
			cDices = rollDice()
			//cDices = []int{1,2,3}
			//cDices = []int{4,2,3}
			cR, _ := getPayout(cDices)
			fmt.Print("cR: ")
			fmt.Println(cR)
			fmt.Println(cDices)
			if cR != 0 {
				break
			}
		}

		winner := decideWinner(pDices, cDices)

		pPayout, _ := getPayout(pDices)
		cPayout, _ := getPayout(cDices)
	
		payout := selectPayout(winner,pPayout,cPayout)

		fmt.Println(payout)
		if winner == "parent" {
			fmt.Println("winner: parent")
			parentMoney += dMoeny * payout
			childMoney -= dMoeny * payout
		} else {
			fmt.Println("winner: child")
			childMoney += dMoeny * payout
			parentMoney -= dMoeny * payout
		}
	}
	pR := float64(parentMoney) / float64(dMoeny * dLoopN)
	cR := float64(childMoney)  / float64(dMoeny * dLoopN)

	return pR , cR
}

//TODO 倍付け
func selectPayout(winner string, pPayout int, cPayout int) int {
	if winner == "parent" {
		if pPayout == -2 {
			return -2
		}

		if cPayout == -2 {
			return 2
		}
		return pPayout
	}

	if pPayout == -2 {
		return 2
	}
	return cPayout
}

func main() {
	rand.Seed(time.Now().UnixNano())
	parentMoney, childMoney := playGame()
	fmt.Printf("Parent Money: %f\n", parentMoney)
	fmt.Printf("Child Money: %f\n", childMoney)
}
