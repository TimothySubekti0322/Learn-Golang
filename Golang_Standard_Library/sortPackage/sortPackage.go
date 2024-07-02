package sortPackage

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i , j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func MainSortPackage() {
	users := []User{
		{"wick", 23},
		{"ethan", 40},
		{"bourne", 35},
		{"wick", 55},
	}

	fmt.Println("----- Before sort -----")
	for _, user := range users {
		fmt.Println(user)
	}

	fmt.Println("----- After sort -----")
	sort.Sort(UserSlice(users))
	for _, user := range users {
		fmt.Println(user)
	}
}