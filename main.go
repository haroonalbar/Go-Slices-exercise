package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var val [][]uint8
	for i:= 0; i< dy; i++{
		b := make([]uint8, dx)
		for j:= 0; j < dx ; j++{
			b[j] = uint8((i + j) /2)
		}
		val = append(val,b)
	}
	return val
}

func main() {
	pic.Show(Pic)
}

