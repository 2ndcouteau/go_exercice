package main

import (
	"fmt"
	"os"
)
	// "bufio"

//flag.Parse()  // scan argument list with option
	// Some example for init flag
		// wordPtr := flag.String("word", "foo", "a string")
		// numbPtr := flag.Int("numb", 42, "an int")
		// boolPtr := flag.Bool("fork", false, "a bool")
		// var svar string
		//     flag.StringVar(&svar, "svar", "bar", "a string var")


func print(str *string) bool{
	fmt.Println(*str)
	return false
}

func main() {
	var str string = "Hello Brasil"

	args := os.Args[1:]
	len := len(args)

	fmt.Println("len args[] = ", len)
	for i := 0 ; i < len ; i++ {
//		len = len([]rune((*args)[i]))
//		fmt.Println("len ", args[i], "= ", len)
		fmt.Println("arg[", i ,"] = ", args[i])
	}
	if args[0] == "quit" {
		fmt.Println("byebye")
		os.Exit(24)
	}
	ret := print(&str)
	if ret {
		fmt.Println("ret == true")
	} else {
		fmt.Println("ret == false")
	}
		// if text != "quit" {
		// 	fmt.Println("Your text was: ", text)
		//}
	os.Exit(42)
}
