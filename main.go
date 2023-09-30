package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const HIFUMI = -2
const ME_NASHI = 0
const ME_NOMAL = 1
const SHIGORO = 2
const ARASHI = 3
const PIN_ZORO = 5

func rollDice() []int {
	dice := []int{rand.Intn(6) + 1, rand.Intn(6) + 1, rand.Intn(6) + 1}
	sort.Ints(dice)
	return dice
}

func diceType(roll []int) (int, int) {
	// fmt.Println(roll)
	// 111
	if roll[0] == 1 && roll[1] == 1 && roll[2] == 1 {
		return PIN_ZORO, 0
	}
	// 456
	if roll[0] == 4 && roll[1] == 5 && roll[2] == 6 {
		return SHIGORO, 0
	}
	// 222,333,444,555,666
	if roll[0] == roll[1] && roll[1] == roll[2] {
		return ARASHI, 0
	}

	// 123
	if roll[0] == 1 && roll[1] == 2 && roll[2] == 3 {
		return HIFUMI, 0
	}

	// 112,221,...
	if roll[0] == roll[1] {
		return ME_NOMAL, roll[2]
	}
	if roll[1] == roll[2] {
		return ME_NOMAL, roll[0]
	}
	if roll[0] == roll[2] {
		return ME_NOMAL, roll[1]
	}
	// 134,...
	return ME_NASHI, 0
}

func decideWinner(parentRoll, childRoll []int) string {
	parentPayout, pDiceN := diceType(parentRoll)
	childPayout, cDiceN := diceType(childRoll)

	if parentPayout == 1 && childPayout == 1 {
		if pDiceN >= cDiceN {
			return "parent"
		}
		return "child"
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
	dLoopN := 100000
	parentMoney := dMoeny
	childMoney := dMoeny

	for i := 0; i < dLoopN; i++ {
		pDices := []int{}
		for j := 0; j < 3; j++ {
			pDices = rollDice()
			//pDices = []int{2,3,3}
			//pDices = []int{3,3,3}
			pR, _ := diceType(pDices)
			fmt.Print("pR: ")
			fmt.Println(pR)
			fmt.Println(pDices)
			if pR != ME_NASHI {
				break
			}
		}

		cDices := []int{}
		for j := 0; j < 3; j++ {
			cDices = rollDice()
			//cDices = []int{2,2,3}
			//cDices = []int{4,2,3}
			cR, _ := diceType(cDices)
			fmt.Print("cR: ")
			fmt.Println(cR)
			fmt.Println(cDices)
			if cR != ME_NASHI {
				break
			}
		}

		winner := decideWinner(pDices, cDices)

		pPayout, _ := diceType(pDices)
		cPayout, _ := diceType(cDices)
	
		payout := payoutMultiplier(winner,pPayout,cPayout)

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


func payoutMultiplier(winner string, pPayout int, cPayout int) int {
	if pPayout == 0 {
		pPayout = 1
	}
	if cPayout == 0 {
		cPayout = 1
	}
	if winner == "parent" {
		if pPayout == HIFUMI {
			baizuke := 1
			if cPayout == ARASHI || cPayout == SHIGORO || cPayout == PIN_ZORO {
				baizuke = cPayout
			}

			return HIFUMI * baizuke
		}

		if cPayout ==  HIFUMI {
			baizuke := 1
			if pPayout == ARASHI || pPayout == SHIGORO || pPayout == PIN_ZORO {
				baizuke = pPayout
			}
			return -HIFUMI * baizuke
		}
		return pPayout
	}

	if pPayout == HIFUMI {
		baizuke := 1
		if cPayout == ARASHI || cPayout == SHIGORO || cPayout == PIN_ZORO {
			baizuke = cPayout
		}
		return -HIFUMI * baizuke
	}
	return cPayout
}

func main() {
	rand.Seed(time.Now().UnixNano())
	parentMoney, childMoney := playGame()
	fmt.Printf("Parent Money: %f\n", parentMoney)
	fmt.Printf("Child Money: %f\n", childMoney)
}
