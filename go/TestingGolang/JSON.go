package main

import (
	"encoding/json"
	"fmt"
)

func main() {
byt := []byte(`{"name":"Alexander","age":18,"isBlocked":true}`)
var dat User
if err := json.Unmarshal(byt, &dat); err != nil {
	panic(err)
}
fmt.Println(dat.Name)
}

/////////////////////////////JSON.MARSHAL///////////////////////
type User struct {
	Name string `json:"name"`
	Age int      `json:"age"`
	IsBlocked bool `json:"isBlocked"`
}

//func main() {
	//sv := map[string]interface{}{"field1": true, "field2": 9, "arr": []string{"1","2","3"}}
//	sv := User{
//    Name: "Alexander",
//	Age: 18,
//	IsBlocked: true,
//	}
//	boolVar, _ := json.Marshal(sv)
//	fmt.Println(string(boolVar))
//}