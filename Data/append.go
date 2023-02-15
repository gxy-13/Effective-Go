package main

func main() {

	x := []int{1, 2, 3}
	x = append(x, 4, 5, 6)
	y := []int{7, 8, 9}
	// 如果没有... 就是编译错误， x需要追加int类型， 但是y是[]int
	x = append(x, y...)

}
