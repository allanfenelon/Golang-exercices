package main

import (
	"fmt"
)

func screen(pt string) (string, string) {
	return pt, "é macaco"
}
func looping(name string) {
	for i := 0; i < 9; i++ {
		fmt.Println(screen(name))
	}
}
func main() {
	var name string
	fmt.Println("Digite um nome:")
	fmt.Scan(&name)
	looping(name)

}
