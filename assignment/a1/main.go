package main

import (
	"fmt"
	"os"
	"strconv"
)


type BioUser struct {
	name string
	job string
	address string
}

type UserDir struct {
	Dir []BioUser
}

func main() {
	var userDir UserDir
        newUser := newUser()
	userDir.PreInsert(newUser)

        noUser := noUserArg(len(userDir.Dir))		
	userDir.PrintUser(noUser)
}

func (u *UserDir) PreInsert(data []BioUser){
	fmt.Println("User Data Pre-Inserted")
	// fmt.Println(data)
	
	for _, value := range data {
		(*u).Dir = append((*u).Dir, value)
	}
}

func noUserArg(userDirLen int) int{
	lenArgs := len(os.Args)
	if lenArgs < 2 {
		fmt.Println("Need Argument(number) For User Index Number to Print The User Data")
		os.Exit(1)
	}
	
        noUser, err := strconv.Atoi(os.Args[1])
        if err != nil {
                fmt.Println(err)
		os.Exit(1)
        }

	if userDirLen < noUser {
		fmt.Println("Index out of range")
                os.Exit(1)
	}

	return noUser
}

func newUser() []BioUser{
	
	return []BioUser{{"Arya","Programmer Php","Depok"}, 
			 {"Rangga","Programmer Go","Depok"},
			 {"Kusuma","Programmer Javascript","Depok"}}	

}

func (u *UserDir) PrintUser(noUser int){
	fmt.Println("No\t :", noUser)
	fmt.Println("Nama\t :", (*u).Dir[noUser].name)
 	fmt.Println("Job\t :", (*u).Dir[noUser].job)
	fmt.Println("Address\t :", (*u).Dir[noUser].address)
}
