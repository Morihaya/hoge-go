package main

// Refer: https://gist.github.com/tetsuok/2280162

import "code.google.com/p/go-tour/pic"

// Sample 1
func Pic(dx, dy int) [][]uint8 {
	// allocate
	p := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		p[i] = make([]uint8, dx)
	}
	// drawing
	for y, _ := range p {
		for x, _ := range p[y] {
			// Choice one of three
			p[y][x] = uint8(x ^ y)
			// p[y][x] = uint8(x * y)
			// p[y][x] = uint8((x * y)/2)
		}
	}
	return p
}

// Sample 2 . これが一番美しいかも
func Pic(dx, dy int) [][]uint8 {

	pic := make([][]uint8, dy)

	for y := range pic {

		pic[y] = make([]uint8, dx)

		for x := range pic[y] {
			pic[y][x] = uint8((x + y) / 2)
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
