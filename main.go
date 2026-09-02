// Go akan menjadikannya executable ketika terdapat func main()
package main

// Kita menggunakan package standard library Go bernama fmt.
// fmt digunakan untuk formatting dan input/output.
import "fmt"

var name = "Willy"

func main() {
	fmt.Println("Hello, " + name + "!")
	fmt.Println("Learning Go Day 01")
	fmt.Println("Built with Go")
}
