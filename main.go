package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Yaku int

const (
	HIFUMI Yaku = iota
	ME_NASHI
	ME
	SHIGORO
	ARASHI
	PIN_ZORO
)

type Dice struct {
	value int
}

func (d *Dice) roll()  {
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

func (d *Dices) Yaku() Yaku{
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
	if first == second || second == third  || first == third { 
		return ME
	}
	return ME_NASHI
}


type Player struct {
	Yaku
	Me int
}

func (p *Player) Roll() {
	for range(3) {
		d := Dices{}
		d.roll()
		if d.Yaku() != ME_NASHI {
			p.Yaku = d.Yaku()
			if p.Yaku == ME {
				p.Me = d.Dices[2].value
			}
			return
		}
	}
	p.Yaku = ME_NASHI
}

type Parent struct {
	Yaku
	Me int
}

func (p *Parent) Roll() {
	for range(3) {
		d := Dices{}
		d.roll()
		if d.Yaku() != ME_NASHI {
			p.Yaku = d.Yaku()
			if p.Yaku == ME {
				p.Me = d.Dices[2].value
			}
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

func (g *Game) Start(ko Player, pa Parent, n int) {
	g.Player = ko
	g.Parent = pa
	g.Player.Roll()
	g.Parent.Roll()
}

func (g Game) Judge() {

}




func main() {
	fmt.Println("Hello, World!")
	rand.NewSource(time.Now().UnixNano())
}

