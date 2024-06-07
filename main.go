package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	for _, argument := range arguments {
		for _, arg := range argument { //prints everyrhing that is passed on commandline argument
			z01.PrintRune(arg)
		}

	}
	z01.PrintRune('\n')
	//prints first argument(firs param)
	for _, first := range arguments[0] {
		z01.PrintRune(first)

	}
	z01.PrintRune('\n')
	//paramcount
	count := len(arguments)
	fmt.Println(len(arguments))
	s := strconv.Itoa(count)
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')

}
