package sort

//insertion sort

func InsertionSort(arr []int) []int {

	n := len(arr)

	var key int
	var j int

	for i := 1; i < n; i++ {

		key = arr[i]

		j = i - 1

		for j >= 0 && arr[j] > key {

			arr[j+1] = arr[j]

			j--

		}

		arr[j+1] = key

	}

	return arr

}
