package gobenchmark

import (
	"cmp"
	"slices"
	"sort"
)

type MyStruct struct {
	Val int
}

//
// non stable parts
//

func SliceNonStableInt(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

func SortNonStableInt(arr []int) {
	slices.SortFunc(arr, func(i, j int) int {
		return cmp.Compare(i, j)
	})
}

func SliceNonStableStruct(arr []MyStruct) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Val < arr[j].Val
	})
}

func SortNonStableStruct(arr []MyStruct) {
	slices.SortFunc(arr, func(i, j MyStruct) int {
		return cmp.Compare(i.Val, j.Val)
	})
}

func SliceNonStablePtr(arr []*MyStruct) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Val < arr[j].Val
	})
}

func SortNonStablePtr(arr []*MyStruct) {
	slices.SortFunc(arr, func(i, j *MyStruct) int {
		return cmp.Compare(i.Val, j.Val)
	})
}

//
// stable parts
//

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
