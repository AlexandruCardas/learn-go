package sorting

func sweep(list []int) {
	for i := 0; i < len(list)-1; i++ {
		first := list[i]
		second := list[i+1]

		if list[i] > list[i+1] {
			list[i] = second
			list[i+1] = first
		}
	}
}

func BubbleSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		sweep(list)
	}

	return list
}
