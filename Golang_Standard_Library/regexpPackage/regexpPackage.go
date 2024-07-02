package regexpPackage

import (
	"fmt"
	"regexp"
)

func MainRegexpPackage() {
	expression := `[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}`
	var regex = regexp.MustCompile(expression)

	fmt.Println(regex.MatchString("testgmail.com"))
	fmt.Println(regex.MatchString("test@gmail.com"))
	fmt.Println(regex.MatchString("test@gmail.c@om"))

	fmt.Println(regex.FindAllString("testgmail.com test@gmail.com test@gmail.c@om", 3))
}