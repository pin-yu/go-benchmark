package gobenchmark

import (
	"cmp"
	"slices"
	"sort"
)

func SliceStableInt(arr []int) {
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

func SortStableInt(arr []int) {
	slices.SortStableFunc(arr, func(i, j int) int {
		return cmp.Compare(i, j)
	})
}

type MyStruct struct {
	Val int
}

func SliceStableStruct(arr []MyStruct) {
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i].Val < arr[j].Val
	})
}

func SortStableStruct(arr []MyStruct) {
	slices.SortStableFunc(arr, func(i, j MyStruct) int {
		return cmp.Compare(i.Val, j.Val)
	})
}

func SliceStablePtr(arr []*MyStruct) {
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i].Val < arr[j].Val
	})
}

func SortStablePtr(arr []*MyStruct) {
	slices.SortStableFunc(arr, func(i, j *MyStruct) int {
		return cmp.Compare(i.Val, j.Val)
	})
}
