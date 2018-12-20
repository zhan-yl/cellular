package cellular

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIslive(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(0, Islive(0, []int{1, 0, 0, 0}), "2*2 test [1,1]")
	ast.Equal(0, Islive(0, []int{1, 1, 0, 0}), "2*2 test [1,1]")
	ast.Equal(1, Islive(0, []int{1, 1, 1, 0}), "2*2 test [1,1]")
	ast.Equal(1, Islive(0, []int{1, 1, 1, 1}), "2*2 test [1,1]")
	ast.Equal(0, Islive(0, []int{0, 0, 0, 0}), "2*2 test [1,1]")
	ast.Equal(0, Islive(0, []int{0, 1, 0, 0}), "2*2 test [1,1]")
	ast.Equal(0, Islive(0, []int{0, 1, 1, 0}), "2*2 test [1,1]")
	ast.Equal(1, Islive(0, []int{0, 1, 1, 1}), "2*2 test [1,1]")

	ast.Equal(0, Islive(1, []int{0, 1, 0, 0, 0, 0, 0, 0, 0}), "3*3 test [1,2]")
	ast.Equal(0, Islive(1, []int{1, 1, 0, 0, 0, 0, 0, 0, 0}), "3*3 test [1,2]")
	ast.Equal(1, Islive(1, []int{1, 1, 1, 0, 0, 0, 0, 0, 0}), "3*3 test [1,2]")
	ast.Equal(1, Islive(1, []int{1, 1, 1, 1, 0, 0, 0, 0, 0}), "3*3 test [1,2]")
	ast.Equal(0, Islive(1, []int{1, 1, 1, 1, 1, 0, 0, 0, 0}), "3*3 test [1,2]")
	ast.Equal(0, Islive(1, []int{1, 1, 1, 1, 1, 1, 0, 0, 0}), "3*3 test [1,2]")
	ast.Equal(0, Islive(1, []int{0, 0, 0, 0, 0, 0, 0, 0, 0}), "3*3 test [1,2]")
	ast.Equal(0, Islive(1, []int{1, 0, 0, 0, 0, 0, 0, 0, 0}), "3*3 test [1,2]")
	ast.Equal(0, Islive(1, []int{1, 0, 1, 0, 0, 0, 0, 0, 0}), "3*3 test [1,2]")
	ast.Equal(1, Islive(1, []int{1, 0, 1, 1, 0, 0, 0, 0, 0}), "3*3 test [1,2]")
	ast.Equal(0, Islive(1, []int{1, 0, 1, 1, 1, 0, 0, 0, 0}), "3*3 test [1,2]")
	ast.Equal(0, Islive(1, []int{1, 0, 1, 1, 1, 1, 0, 0, 0}), "3*3 test [1,2]")
	ast.Equal(0, Islive(2, []int{1, 1, 0, 1, 0, 0, 0, 0, 0}), "3*3 test [1,2]")
}

func TestLiveCell(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([]int{0, 0, 0, 0}, LiveCell([]int{1, 0, 0, 0}), "2*2 cell [1,1]")
	ast.Equal([]int{1, 1, 1, 1}, LiveCell([]int{1, 1, 1, 0}), "2*2 cell [1,1]")
	ast.Equal([]int{1, 1, 0, 1, 1, 0, 0, 0, 0}, LiveCell([]int{1, 1, 0, 1, 0, 0, 0, 0, 0}), "3*3 cell [1,1]")
	ast.Equal([]int{0, 0, 0, 0, 0, 0, 0, 0, 0}, LiveCell([]int{0, 0, 0, 1, 1, 1, 0, 0, 0}), "3*3 cell [1,1]")
	ast.Equal([]int{0, 0, 0, 0, 1, 1, 0, 1, 1}, LiveCell([]int{1, 0, 0, 1, 1, 0, 1, 1, 1}), "3*3 cell [1,1]")

	ast.Equal([]int{1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0}, LiveCell([]int{1, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1}), "3*3 cell [1,1]")
}
