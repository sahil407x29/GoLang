package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("how many stars for pizza:")

	// comma ok || comma err 
	input,_:= reader.ReadString('\n') //'\n' here is passed as an ender (till which point to read)
	// this second param (_) is passed as error catching variable which will store the error incase anything goes wrong
	fmt.Println("thanks for rating ",input)
}