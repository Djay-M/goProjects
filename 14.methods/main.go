package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to methods module")
	djay := Users{Name: "Djay M", Age: 29, Phone: 7411856661, IsArchived: false}
	djay.getName()
	djay.updateUserPhone(1234567890)
	fmt.Printf("user = %+v \n", djay)
}

type Users struct{
	Name string
	Age int
	Phone int64
	IsArchived bool
}

func (user Users) getName() {
	fmt.Printf("User name is: %v \n", user.Name)
}

func (user Users) updateUserPhone(newPhone int) {
	fmt.Printf("old user phone number is %v \n", user.Phone)
	user.Phone = int64(newPhone)
	fmt.Printf("updated user phone is: %v \n", user.Phone)
}