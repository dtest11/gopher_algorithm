package back_tracking

import (
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	input := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	result := combinationSum2(input, target)
	t.Log(result)
}
