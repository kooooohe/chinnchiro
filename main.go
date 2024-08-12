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
}

func (p *Player) Roll() {
	for i := 0; i < 3; i++ {
		d := Dices{}
		d.roll()
	}
	p.Yaku = ME_NASHI
}


func main() {
	fmt.Println("Hello, World!")
	rand.NewSource(time.Now().UnixNano())
}

