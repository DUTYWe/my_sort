package main

import (
	"golang.org/x/exp/constraints"
)

//take the minimum value
func less[T constraints.Ordered](x, y T) bool {
	return x < y
}

//swap elements in an array
func swap[T constraints.Ordered](a []T, i, j int) {
	a[i], a[j] = a[j], a[i]
}

//select sort
func SelectSort[T constraints.Ordered](nums []T) {
	for i := 0; i < len(nums)-1; i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if less(nums[j], nums[min]) {
				min = j
			}
		}
		swap(nums, i, min)
	}
}

//bubble sort
func BubbleSort[T constraints.Ordered](nums []T) {

	for i := 0; i < len(nums); i++ {
		//flag of early exit cycle
		flag := false
		for j := 0; j < len(nums)-i-1; j++ {
			if less(nums[j+1], nums[j]) {
				swap(nums, j+1, j)
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

//insertion sort
func InsertionSort[T constraints.Ordered](nums []T) {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && less(nums[j], nums[j-1]); j-- {
			swap(nums, j-1, j)
		}
	}
}

//shell sort
func ShellSort[T constraints.Ordered](nums []T) {
	shell := 1
	for shell < len(nums)/3 {
		shell = 3*shell + 1
	}
	for shell >= 1 {
		for i := shell; i < len(nums); i++ {
			for j := i; j >= shell && less(nums[j], nums[j-shell]); j -= shell {
				swap(nums, j, j-shell)
			}
		}
		shell /= 3
	}
}

//quick sort
func partition[T constraints.Ordered](nums []T, lo, hi int) int {
	pivot := nums[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if nums[j] < pivot {
			i++
			swap(nums, i, j)
		}
	}
	swap(nums, i, hi)
	return i
}
func QuickSort[T constraints.Ordered](nums []T, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(nums, lo, hi)
	QuickSort(nums, lo, p-1)
	QuickSort(nums, p+1, hi)
}
