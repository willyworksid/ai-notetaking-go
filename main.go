// Go akan menjadikannya executable ketika terdapat func main()
package main

// Kita menggunakan package standard library Go bernama fmt.
// fmt digunakan untuk formatting dan input/output.
import (
	"fmt"

	"github.com/willyworksid/ai-notetaking-go/greeting"
)

func main() {
	fmt.Println(greeting.Hello("Willy"))
	fmt.Println(greeting.secret()) // error: secret is not exported

}
