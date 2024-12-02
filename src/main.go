// Purpose: demonstrate scanning/lexical analysis

package main 

import (
	"os"
	"fmt"
)

func main(){
	argCount := len(os.Args)

	if argCount > 2{
		fmt.Println("Usage: pygo [script]")
		os.Exit(64) // you did not follow correct usage
	} else if argCount == 2 {
		runScript(os.Args[1])
	} else {
		startREPL()
	}
}	

