package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	argslen := len(args)
	if argslen < 2 {
		fmt.Println("rootsh update|check git|emacs")
		os.Exit(0)
	}
	switch args[0] {
	case "update":
		go update(args[1])
	case "check":
		go check(args[1])
	default:
		fmt.Println("You must put some command")
	}
	fmt.Println(args)
	fmt.Println("Hello world")
}

func update(p string) {
	fmt.Printf("update packages %s\n", p)

}

func check(p string) {
	fmt.Printf("Check update %s\n", p)

}
