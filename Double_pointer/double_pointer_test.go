package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var data = map[int][]int{
	9: []int{2, 7, 11, 15},
	//6: []int{3, 2, 4},
	6: []int{3, 3},
}

type d struct {
	Target int
	Ele    []int
	Want   []int
}

func getEle() []d {
	return []d{
		{
			Target: 9,
			Ele:    []int{2, 7, 11, 15},
			Want:   []int{0, 1},
		},
		{
			Target: 6,
			Ele:    []int{3, 2, 4},
			Want:   []int{1, 2},
		},
		{
			Target: 6,
			Ele:    []int{3, 3},
			Want:   []int{0, 1},
		},
		{
			Target: -8,
			Ele:    []int{-1, -2, -3, -4, -5},
			Want:   []int{2, 4},
		},
	}
}

func TestSort(t *testing.T) {
	for _, v := range data {
		sort(v)
		t.Log(v)
	}
}

func TestTWoSum(t *testing.T) {
	for _, v := range getEle() {
		res := twoSum(v.Ele, v.Target)
		assert.Equal(t, v.Want, res, fmt.Sprintf("%+v,Error", v))
	}
}
