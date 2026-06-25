package main

import "fmt"

type Box struct {
	n int
}

func incByValue(b Box) {
	b.n++
}
func incByPointer(b *Box) {
	b.n++
}

func main() {
	box := Box{n: 1}
	incByValue(box)

	fmt.Println(box)

	incByPointer(&box)
	fmt.Println(box)
}