//category: object oriented
//ref: https://opensourcedoc.com/golang-programming/class-object/

package main

import "fmt"

// constructor
func NewUser(name string,job string,age int) *User {
	u := new(User)
	u.name = name
	u.job = job
	u.age = age
	
	return u
}

// data member
type User struct {
	name string
	job string
	age int
}

// function meber
func (u *User) getName() string {
	return u.name
}

func (u *User) getJob() string {
	return u.job
}

func main(){
	yale := NewUser("yale","coder",35)
	fmt.Println(yale.getName())
	
}