package main

import (
	"fmt"
)

func main(){



multiple := func(n1,n2 int)int{
	return n1*n2
}

sumCallBack := func(n1,n2 int) int{
	return n1 + n2 
}

result := doSomething(sumCallBack, "Plus")
 fmt.Println(result)

result = doSomething(multiple, "umnozh")
fmt.Println(result)
}

func doSomething(callBack func(int,int)int, s string) int{
	result:= callBack(5,1)
	fmt.Println(s)
	return result
}


