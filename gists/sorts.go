package main

import (
	"fmt"
)

func bublleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		flag := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
				flag = true
			}

		}
		if !flag {
			break
		}
	}

}

func selectionSort(arr []int) {
	min := 0

	for i := 0; i < len(arr); i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}

		}
		arr[i], arr[min] = arr[min], arr[i]

	}
}

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		val := arr[i]
		j := 0

		flag := false
		for j = i; j > 0; j-- {
			if arr[j-1] <= val {
				break
			}
			arr[j] = arr[j-1]
			flag = true

		}

		if flag {
			arr[j] = val
		}

	}

}

func mergeSort(s []int) []int {
	if len(s) == 1 {
		return s
	}

	lhs := s[:len(s)/2]
	rhs := s[len(s)/2:]

	ret := merge(mergeSort(lhs), mergeSort(rhs))
	return ret

}

func merge(lhs, rhs []int) []int {
	ret := make([]int, 0, len(lhs)+len(rhs))

	for len(lhs) != 0 && len(rhs) != 0 {
		if lhs[0] <= rhs[0] {
			ret = append(ret, lhs[0])
			lhs = lhs[1:]
		} else {
			ret = append(ret, rhs[0])
			rhs = rhs[1:]
		}

	}

	if len(lhs) != 0 {
		ret = append(ret, lhs...)
	}
	if len(rhs) != 0 {
		ret = append(ret, rhs...)
	}

	return ret

}

func shellInsert(arr []int, gap int) {
	for i := 0; i < len(arr); i += gap {

		val := arr[i]
		j := 0

		flag := false
		for j = i; j > 0; j -= gap {
			if arr[j-gap] <= val {
				break
			}
			arr[j] = arr[j-gap]
			flag = true
		}

		if flag {
			arr[j] = val
		}

	}

}

func shellSort(s []int) {
	for gap := len(s) / 2; gap >= 1; gap /= 2 {
		shellInsert(s, gap)
	}

}

func quickSort(s []int, start, end int) {
	if end <= start {
		return
	}

	base := s[start]

	left := start
	right := end

	for left < right {
		for ; left < right && s[right] > base; right-- {
		} //左洞 右开始

		if left < right {
			s[left] = s[right]
			left++
		}

		for ; left < right && s[left] < base; left++ {
		}

		if left < right {
			s[right] = s[left]
			right--
		}

	}
	s[left] = base
	quickSort(s, start, left-1)
	quickSort(s, left+1, end)

}

func adjustHeap(s []int, root int) {
	lhs := root*2 + 1
	rhs := root*2 + 2

	if lhs >= len(s) {
		return
	}

	largest := root

	if s[lhs] > s[largest] {
		largest = lhs
	}

	if rhs < len(s) && s[rhs] > s[largest] {
		largest = rhs
	}

	if largest != root {
		s[root], s[largest] = s[largest], s[root]
		adjustHeap(s, largest)
	}

}

func makeHeap(s []int) {
	for i := len(s)/2 - 1; i >= 0; i-- {
		adjustHeap(s, i)
	}

}

func heapSort(s []int) {
	makeHeap(s)

	for j := len(s) - 1; j >= 1; j-- {
		s[j], s[0] = s[0], s[j]

		adjustHeap(s[:j], 0)
		//fmt.Println(s)
	}

}
func main() {
	arr := []int{3, 2, 1, 4, -1, 30, 50, 1, 0, -2, 100, -200, 101, -5}

	//bublleSort(arr)
	//selectionSort(arr)
	//insertSort(arr)
	//arr = mergeSort(arr)
	//shellSort(arr)
	//quickSort(arr, 0, len(arr)-1)
	heapSort(arr)

	fmt.Println(arr)
}
