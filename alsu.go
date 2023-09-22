package main

import (
	"fmt"
	"os"
)

func main() {

    fmt.Println(len(os.Args), os.Args)
    
	if (len(os.Args) > 1) && (os.Args[1] == "analyse") {
		analyse()
	} else if (len(os.Args) > 1) && (os.Args[1] == "test") {
		test()
	} else {
		fmt.Println("to help press alsu -h")
	}
}

func analyse(){
	if len(os.Args) == 2 { 
		fmt.Println("to help press alsu analyse -h")
		return
	}
	if (os.Args[2] != "-h") {
		fmt.Println("analyse")
	}else{
		fmt.Println("this is help analyse")
	}
}

func test(){
	if len(os.Args) == 2 { 
		fmt.Println("to help press alsu test -h")
		return
	}
	if (os.Args[2] != "-h") {
		fmt.Println("test")
	}else{
		fmt.Println("this is help test")
	}
}
