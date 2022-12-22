package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
)

func main() {
	fmt.Println("Current process id is: ", os.Getegid())
	fmt.Println("Parent process id is: ", os.Getppid())
	fmt.Println("User id: ", os.Geteuid())
	fmt.Println("group id: ", os.Getgid())
	groups, err := os.Getgroups()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("groups ", groups)
	for _, gId := range groups {
		userGroup, err := user.LookupGroupId(strconv.Itoa(gId))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Group Id:", fmt.Sprintf("%4d", gId), " Name:", userGroup.Name)
	}
}
