package main

import "fmt"

func main() {
	var s = []byte{'a', 'b', 'c'}
	s = Append(s, []byte{'d', 'e', 'f'})
	fmt.Println(s)

	// 多维切片
	type LinesOfText [][]byte
	text := LinesOfText{
		[]byte("Now is the time"),
		[]byte("for all good gophers"),
		[]byte("to bring some fun to the party."),
	}
	fmt.Println(text)

	// 在处理像素的扫描行是，必须分配一个二维数组
	// 一种是独立地分配每一个切片，另一种就是只分配一个数组，将各个切片都指向它
	// 若切片会增长或收缩，就应该通过独立分配来避免覆盖下一行，若不会，用单次分配来构造会更加高效
	// 首先是一次一行的：
	// 分配底层切片
	picture := make([][]uint8, 8) // 8行
	// 循环遍历每一行
	for i := range picture {
		picture[i] = make([]uint8, 8) // 每一行是8列
	}

	// 现在是一次分配，对行进行切片
	picture2 := make([][]uint8, 8) // 8行
	// 分配一个大一些的切片以容纳所有的元素
	pixels := make([]uint8, 8*8) // 指定类型uint8。 即便图片是[][]uint8
	// 循环遍历图片所有行，从生育像素切片的前面对每一行进行切片
	for i := range picture2 {
		// 将pixels中的内容按8个一组填充进picture2中， 填8个就切8个
		picture2[i], pixels = pixels[:8], pixels[8:]
	}

}

func Append(slice, data []byte) []byte {
	l := len(slice)
	if l+len(data) > cap(slice) {
		newSlice := make([]byte, (l+len(slice))*2)
		copy(newSlice, slice)
		// len， cap 都赋值给原来的slice
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	return slice
}
