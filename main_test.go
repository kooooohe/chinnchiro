package main

import (
	"testing"
        "github.com/google/go-cmp/cmp"
  	"github.com/stretchr/testify/assert"
)

func newDices(d1, d2, d3 int) Dices {
	return Dices{
		Dices: [3]Dice{
			{value: d1},
			{value: d2},
			{value: d3},
		},
	}
}

func Test_nothing(t *testing.T) {
}

func Test_Dices_Sort(t *testing.T) {
	v := Dices{
		Dices: [3]Dice{
			{value: 1},
			{value: 3},
			{value: 2},
		},
	}
	e := Dices{
		Dices: [3]Dice{
			{1},
			{2},
			{3},
		},
	}
	v.sort()

	opt := cmp.AllowUnexported(Dice{})

	if diff := cmp.Diff(v, e, opt); diff != "" {
		t.Errorf("X value is mismatch (-v +e):%s\n", diff)
	}
}

func Test_Dices_Sort_2(t *testing.T) {
	v := Dices{
		Dices: [3]Dice{
			{value: 3},
			{value: 3},
			{value: 2},
		},
	}
	e := Dices{
		Dices: [3]Dice{
			{2},
			{3},
			{3},
		},
	}
	v.sort()

	opt := cmp.AllowUnexported(Dice{})

	if diff := cmp.Diff(v, e, opt); diff != "" {
		t.Errorf("X value is mismatch (-v +e):%s\n", diff)
	}
}

func Test_Dices_Sort_3(t *testing.T) {
	v := Dices{
		Dices: [3]Dice{
			{value: 3},
			{value: 2},
			{value: 1},
		},
	}
	e := Dices{
		Dices: [3]Dice{
			{1},
			{2},
			{3},
		},
	}
	v.sort()

	opt := cmp.AllowUnexported(Dice{})

	if diff := cmp.Diff(v, e, opt); diff != "" {
		t.Errorf("X value is mismatch (-v +e):%s\n", diff)
	}
}

func Test_Dices_Sort_4(t *testing.T) {
	v := Dices{
		Dices: [3]Dice{
			{value: 1},
			{value: 2},
			{value: 1},
		},
	}
	e := Dices{
		Dices: [3]Dice{
			{1},
			{1},
			{2},
		},
	}
	v.sort()

	opt := cmp.AllowUnexported(Dice{})

	if diff := cmp.Diff(v, e, opt); diff != "" {
		t.Errorf("X value is mismatch (-v +e):%s\n", diff)
	}
}
func Test_Dices_Sort_5(t *testing.T) {
	v := Dices{
		Dices: [3]Dice{
			{value: 1},
			{value: 2},
			{value: 3},
		},
	}
	e := Dices{
		Dices: [3]Dice{
			{1},
			{2},
			{3},
		},
	}
	v.sort()

	opt := cmp.AllowUnexported(Dice{})

	if diff := cmp.Diff(v, e, opt); diff != "" {
		t.Errorf("X value is mismatch (-v +e):%s\n", diff)
	}
}

func Test_Yaku(t *testing.T) {
	ts := []Dices{}
	ns := [][]int{
		{1,1,1},
		{3,1,2},
		{5,4,6},
		{2,2,2},
		{1,1,2},
		{1,4,5},
	}
	es := []Yaku{
		PIN_ZORO,
		HIFUMI,
		SHIGORO,
		ARASHI,
		ME,
		ME_NASHI,
	}
	for _, n := range ns {
		d := newDices(n[0], n[1], n[2])
		d.sort()
		ts = append(ts, d)
	}
	for i, tt := range ts {
		assert.Equal(t,tt.Yaku(), es[i])
	}
}
