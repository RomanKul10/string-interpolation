package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User
	user.UserName = ReadString("What is your name?")
	user.Age = ReadInt("How old are you?")
	user.FavouriteNumber = ReadFloat("What is your favourite number?")

	//fmt.Println("Your name is "+userName+".You are", age, "years old.")
	//fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old.", userName, age))
	fmt.Printf("Your name is %s, you are %d years old. You are favourite number %.4
	f.\n",
		user.UserName,
		user.Age,
		user.FavouriteNumber)
}
func prompt() {
	fmt.Print("->")
}
func ReadString(s string) string {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("Please enter value")

		} else {
			return userInput
		}
	}
}
func ReadInt(s string) int {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}

}
func ReadFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}

}
