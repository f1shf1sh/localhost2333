package main

import (
	"fmt"
	"html/template"
	"vm"
)

func main() {
	// d0g3{Go1aN9_vM_1S_VERY_e@$Y!!}
	var input string
	var retStr template.HTML
	fmt.Println("Please input: ")
	fmt.Scanf("%s", &input)
	retStr = vm.MainVM(input)
	fmt.Println(retStr)
}
