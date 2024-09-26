package main
import "fmt"
func main() {
	fmt.Println("Hello, world!")
	Greeting()
	advanceGreeting("Amaan")
}
func Greeting(){
	fmt.Println("Hey there!")	
}

func advanceGreeting(name string) {
	fmt.Printf("Hello, %s!\n", name)
}