package array

func findDiagonalOrder(nums [][]int) []int {
	//[0, 0]
	//[0, 1], [1, 0]
	//[0, 2], [1, 1], [2, 0]
	//[0, 4], [3, 1]

	//fist - create [][]int and put diagonal elements there
	//then flatten it
	var buff [][]int
	for i, arr := range nums {
		for j, value := range arr {
			if i+j == len(buff) {
				buff = append(buff, []int{})
			}
			buff[i+j] = append(buff[i+j], value)
		}
	}
	var res []int
	for _, arr := range buff {
		for i := len(arr) - 1; i >= 0; i-- {
			res = append(res, arr[i])
		}
	}
	return res
}
