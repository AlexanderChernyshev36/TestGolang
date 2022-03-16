package main

import "fmt"

type structHere struct{
	 N1,N2 int
}

func(s *structHere) sum() int{
return s.N1 + s.N2
}

type interfaceHere interface{
  sum() int
}

func main(){
var a interfaceHere
sh := structHere{1,2}
os := otherStruct{2,3}
a = &sh
fmt.Println(a.sum())
a = &os
fmt.Println(a.sum())
var i interfaceHere = otherStruct{4,5}
fmt.Println(i.sum())
}

type otherStruct struct{
	A,B int 
}

func(o otherStruct) sum() int{
	return o.A + o.B
}