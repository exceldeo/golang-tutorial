package main

import (
	"bufio"
	"fmt"
	"os"
	"user-socialmedia/user"
)

func main() {
	users := user.NewUsers()
	var input int
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Activity Social Media: \n1. register new user\n2. setup\n3. add new post\n4. show all users\n5. exit\n")
		fmt.Scanln(&input)
		switch input {
		case 1:
			fmt.Println("Input your name: ")
			scanner.Scan()
			usr := user.NewUser(scanner.Text())
			users.AddUsers(usr.(*user.User))
		case 2:
			fmt.Println("Setup")
			scanner.Scan()
		case 3:
			fmt.Println("Enter name of user: ")
			scanner.Scan()
		case 4:
			fmt.Println("Show all users : ")
			for _, v := range users.Users {
				fmt.Println("- ", v.Name)
			}
		case 5:
			fmt.Println("Exit")
			os.Exit(0)
		default:
			fmt.Println("Invalid input")
		}
	}
}
