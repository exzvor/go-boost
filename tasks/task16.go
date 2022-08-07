package tasks

/*
 Задача 16. Реализовать быструю сортировку массива (quicksort) встроенными методами
языка.
*/

import "fmt"

func Partition(arr []int, low, high int) int {
	index := low - 1
	pivotElement := arr[high]
	for i := low; i < high; i++ {
		if arr[i] <= pivotElement {
			index += 1
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	arr[index+1], arr[high] = arr[high], arr[index+1]
	return index + 1
}

// QuicksortRange сортирует массив
func QuicksortRange(arr []int, low, high int) {
	if len(arr) <= 1 {
		return
	}

	if low < high {
		pivot := Partition(arr, low, high)
		QuicksortRange(arr, low, pivot-1)
		QuicksortRange(arr, pivot+1, high)
	}
}

// TaskSixteen сортирует весь массив
func TaskSixteen(arr []int) []int {
	QuicksortRange(arr, 0, len(arr)-1)
	fmt.Println(arr)
	return arr
}
