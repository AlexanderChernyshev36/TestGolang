package main

import "fmt"

var _ User = &user{}

type user struct{
	FIO, Adress, Phone string
}

func (u *user) ChangeFIO(newFio string){
	u.FIO = newFio
}

func (u *user) ChangeAdress(newAdress string){
	u.Adress = newAdress
}

type User interface{
  ChangeFIO(newFIO string)
  ChangeAdress(newAdress string)
}

 func NewUser(fio,address,phone string) User{
	 u := user{
		 FIO: fio,
		Adress: address,
		Phone: phone}
		return &u
 }

func main(){
user1 := user{
	FIO: "Chernyshev",
Adress:"Voronezh",
Phone:"8952010101",
}
fmt.Println(user1)
}