package cellular

import (
	"math"
)

// 是否是邻居
func isNeighbor(p int, location int, size int) bool {
	prow := p / size
	pcol := p % size
	lrow := location / size
	lcol := location % size

	if (prow-1 == lrow) && (pcol+1 == lcol || pcol == lcol || pcol-1 == lcol) {
		return true
	} else if (prow == lrow) && (pcol+1 == lcol || pcol-1 == lcol) {
		return true
	} else if (prow+1 == lrow) && (pcol+1 == lcol || pcol == lcol || pcol-1 == lcol) {
		return true
	}

	return false
}

// 根据本身的状态及邻居的状态确定下一轮的状态
func Islive(p int, a []int) int {
	size := int(math.Sqrt(float64(len(a))))

	var sum int
	var state int
	for i, s := range a {
		if i == p {
			state = s
		} else if isNeighbor(p, i, size) {
			sum += s
		}
	}
	if ((state == 1) && (sum == 2 || sum == 3)) || (state == 0 && sum == 3) {
		return 1
	} else {
		return 0
	}
}

// 细胞迭代状态
func LiveCell(in []int) []int {
	out := make([]int, len(in))
	for {
		changed := 0
		for i, _ := range in {
			out[i] = Islive(i, in)
			if out[i] != in[i] {
				changed = 1
			}
		}
		in = out
		if changed == 0 {
			break
		}
	}
	return out
}
