package main

import "fmt"

func main() {
var a interface{} = true
switch a.(type) {
case string:
	fmt.Println("a is a string")
case int:
	fmt.Println("a is a int")
default:
	fmt.Printf("NON type %T", a)
}
}