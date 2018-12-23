package cellular

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const T = true
const F = false

func TestField(t *testing.T) {
	ast := assert.New(t)

	f := &Field{s: [][]bool{
		{F, T, F},
		{F, T, F},
		{F, T, F},
	}, w: 3, h: 3}

	/*
		f := NewField(3, 3)
		f.Set(1, 0, T)
		f.Set(1, 1, T)
		f.Set(1, 2, T)
	*/

	ast.Equalf(T, f.Alive(1, 0), "x[%d][%d] is alive", 1, 0)
	ast.Equalf(T, f.Alive(1, 1), "x[%d][%d] is alive", 1, 1)
	ast.Equalf(T, f.Alive(1, 2), "x[%d][%d] is alive", 1, 2)

	ast.Equalf(T, f.Next(1, 0), "x[%d][%d] next step is alive", 1, 0)
	ast.Equalf(T, f.Next(1, 1), "x[%d][%d] next step is alive", 1, 1)
	ast.Equalf(T, f.Next(1, 2), "x[%d][%d] next step is alive", 1, 2)
	ast.Equalf(T, f.Next(2, 1), "x[%d][%d] next step is alive", 2, 1)
	ast.Equalf(T, f.Next(1, 0), "x[%d][%d] next step is alive", 1, 0)
	ast.Equalf(T, f.Next(0, 1), "x[%d][%d] next step is alive", 0, 1)
	ast.Equalf(T, f.Next(2, 0), "x[%d][%d] next step is alive", 2, 0)
	ast.Equalf(T, f.Next(2, 2), "x[%d][%d] next step is alive", 2, 2)
}

func TestLife(t *testing.T) {
	ast := assert.New(t)

	life := &Life{
		a: &Field{s: [][]bool{
			{F, T, F},
			{F, T, F},
			{F, T, F},
		}, w: 3, h: 3},
		b: NewField(3, 3),
		w: 3, h: 3,
	}
	/*
		life.a.Set(0, 0, T)
		life.a.Set(0, 1, T)
		life.a.Set(1, 0, T)
		ast.Equalf(T, life.a.Alive(0, 2), "x[%d][%d] is alive", 0, 2)
		ast.Equalf(T, life.a.Alive(1, 2), "x[%d][%d] is alive", 1, 2)
	*/
	life.Step()
	ast.Equalf(&Field{s: [][]bool{
		{T, T, T},
		{T, T, T},
		{T, T, T},
	}, w: 3, h: 3}, life.a, "life[3X3] step one")
	life.Step()
	ast.Equalf(&Field{s: [][]bool{
		{F, F, F},
		{F, F, F},
		{F, F, F},
	}, w: 3, h: 3}, life.a, "life[3X3] step two")
	life.Step()
	ast.Equalf(&Field{s: [][]bool{
		{F, F, F},
		{F, F, F},
		{F, F, F},
	}, w: 3, h: 3}, life.a, "life[3X3] step tree")

	life = &Life{
		a: &Field{s: [][]bool{
			{F, F, T, F, F},
			{F, F, T, F, F},
			{T, T, T, T, T},
			{F, F, T, F, F},
			{F, F, T, F, F},
		}, w: 5, h: 5},
		b: NewField(5, 5),
		w: 5, h: 5,
	}

	life.Step()
	ast.Equalf(&Field{s: [][]bool{
		{F, T, T, T, F},
		{T, F, F, F, T},
		{T, F, F, F, T},
		{T, F, F, F, T},
		{F, T, T, T, F},
	}, w: 5, h: 5}, life.a, "life[5X5] step one")
	life.Step()
	ast.Equalf(&Field{s: [][]bool{
		{F, F, F, F, F},
		{F, F, T, F, F},
		{F, T, F, T, F},
		{F, F, T, F, F},
		{F, F, F, F, F},
	}, w: 5, h: 5}, life.a, "life[5X5] step two")

	life = NewLife(10, 10)
	fmt.Print("\x0c", life)
}
