package main

import (
	"fmt"

	"github.com/vagababov/litelibs/algo/unionfind"
)

func main() {
	uf := unionfind.NewQU(10)
	for {
		fmt.Print("Enter components to connect: ")
		var p, q int
		if n, err := fmt.Scan(&p, &q); err != nil || n != 2 {
			fmt.Println("done!")
			break
		}
		uf.Union(p, q)
	}
	fmt.Printf("\n%v\n", uf)
}
