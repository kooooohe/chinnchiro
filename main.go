package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Yaku int

const (
	HIFUMI   Yaku = 1
	ME_NASHI Yaku = 2
	ME       Yaku = 3
	SHIGORO  Yaku = 4
	ARASHI   Yaku = 5
	PIN_ZORO Yaku = 6
)

type Dice struct {
	value int
}

func (d *Dice) roll() {
	d.value = rand.Intn(6) + 1
}

type Dices struct {
	Dices [3]Dice
}

func (d *Dices) roll() {
	for i := 0; i < 3; i++ {
		d.Dices[i].roll()
	}
	d.sort()
}

func (d *Dices) sort() {
	for i := 0; i < 3; i++ {
		for j := i; j < 3; j++ {
			if d.Dices[i].value > d.Dices[j].value {
				d.Dices[i], d.Dices[j] = d.Dices[j], d.Dices[i]
			}
		}
	}
}

func (d *Dices) Yaku() Yaku {
	first := d.Dices[0].value
	second := d.Dices[1].value
	third := d.Dices[2].value

	if first == 1 && second == 1 && third == 1 {
		return PIN_ZORO
	}
	if first == second && second == third {
		return ARASHI
	}
	if first == 4 && second == 5 && third == 6 {
		return SHIGORO
	}
	if first == 1 && second == 2 && third == 3 {
		return HIFUMI
	}
	if first == second || second == third || first == third {
		return ME
	}
	return ME_NASHI
}

type Player struct {
	Yaku
	Me int
	Dices
}

func (p *Player) Roll() {
	for range 3 {
		d := Dices{}
		d.roll()
		if d.Yaku() != ME_NASHI {
			p.Yaku = d.Yaku()
			if p.Yaku == ME {
				if d.Dices[0].value == d.Dices[1].value {
					p.Me = d.Dices[2].value
				} else {
					p.Me = d.Dices[0].value
				}
			}
			p.Dices = d
			return
		}
	}
	p.Yaku = ME_NASHI
}

type Parent struct {
	Yaku
	Me int
	Dices
}

func (p *Parent) Roll() {
	for range 1000000 {
		d := Dices{}
		d.roll()
		if d.Yaku() != ME_NASHI {
			p.Yaku = d.Yaku()
			if p.Yaku == ME {
				if d.Dices[0].value == d.Dices[1].value {
					p.Me = d.Dices[2].value
				} else {
					p.Me = d.Dices[0].value
				}
			}
			p.Dices = d
			return
		}
	}
	p.Yaku = ME_NASHI
}

type Game struct {
	Bet int
	Player
	Parent
}

func (g *Game) Start(n int) {
	m := map[Yaku]string{
		1: "HIFUMI",
		2: "ME_NASHI",
		3: "ME",
		4: "SHIGORO",
		5: "ARASHI",
		6: "PIN_ZORO",
	}
	_ = m

	pc := 0
	cc := 0
	for range n {
		g.Player = Player{}
		g.Parent = Parent{}
		g.Player.Roll()
		g.Parent.Roll()
		isPwin, r := g.Judge()
		/*
			fmt.Println("parent:", g.Parent)
			fmt.Println(m[g.Parent.Yaku])
			fmt.Println("children:", g.Player)
			fmt.Println(m[g.Player.Yaku])
			fmt.Println(g.Judge())
			fmt.Println()
		*/
		if isPwin {
			pc += r
			//fmt.Println("parent:", r)
		} else {
			cc += r
			//fmt.Println("children:", r)
		}
	}
	fmt.Println("parent:", pc)
	fmt.Println("player:", cc)
	fmt.Println("P:", float64(pc)/float64(cc))
}

func (g Game) Judge() (isParentWin bool, baizuke int) {
	baizuke = 1
	if g.isParentWin() {
	} else {
		if g.Parent.Yaku == HIFUMI {
			baizuke *= 2
		}
		if g.Player.Yaku == PIN_ZORO {
			baizuke *= 3
		}
		if g.Player.Yaku == ARASHI {
			baizuke *= 2
		}
		if g.Player.Yaku == SHIGORO {
			baizuke *= 2
		}
	}

	return g.isParentWin(), baizuke
}
func (g Game) isParentWin() bool {
	if g.Parent.Yaku == ME && g.Player.Yaku == ME {
		return g.Parent.Me >= g.Player.Me
	}
	return g.Parent.Yaku >= g.Player.Yaku
}

func main() {
	rand.NewSource(time.Now().UnixNano())
	g := Game{}
	g.Start(10_000_000)
}
