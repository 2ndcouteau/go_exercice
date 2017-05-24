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


func print(str *string) bool {
	fmt.Print(*str + "\n")
	return false
}

func main() {
	var str string = "Hello Brasil"

	args := os.Args[1:]
	len := len(args)

	fmt.Print("len args[] = ", len, "\n")
	for i := 0 ; i < len ; i++ {
//		len([]rune(args[i]))
//		fmt.Print("len ", args[i], "= ", len)
	}
	if args[0] == "quit" {
		fmt.Print("byebye\n")
		return
	}
	ret := print(&str)
	if ret {
		fmt.Print("ret == true\n")
	} else {
		fmt.Print("ret == false\n")
	}
		// if text != "quit" {
		// 	fmt.Println("Your text was: ", text)
		//}
}
