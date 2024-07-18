package main

import (
	"bufio" // buffered input
	"os"
	// "io"
	"fmt"
)

func main(){
	
	fmt.Println("enter text blob to throw :")
	reader := bufio.NewReader(os.Stdin)  // catch user input into a 
	line, _ := reader.ReadString('\n')
	// if err == Nil:

	fmt.Println("i saw this", line)


	var cow = `        \  ^__^
         \ (oo)\_______
	   (__)\       )\/\
	       ||----w |
	       ||     ||
	`
	fmt.Println(cow)
}