package main   // Package declaration; 
			// every Go progeam should start with paackage declaration
import "fmt"
func main() {
	for i:= 1; i<= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is Even")
		} else {
			fmt.Println(i, "is Odd")
		}
	}
}
