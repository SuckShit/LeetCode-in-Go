package problem1074

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	matrix [][]int
	target int
	ans    int
}{

	{
		[][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}},
		0,
		4,
	},

	{
		[][]int{{1, -1}, {-1, 1}},
		0,
		5,
	},

	// 可以有多个 testcase
}

func Test_numSubmatrixSumTarget(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, numSubmatrixSumTarget(tc.matrix, tc.target), "输入:%v", tc)
	}
}

func Benchmark_numSubmatrixSumTarget(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numSubmatrixSumTarget(tc.matrix, tc.target)
		}
	}
}
