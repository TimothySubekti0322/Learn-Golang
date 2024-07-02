package main

import "fmt"

func main(){
	counter := 1

	for counter <= 3 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	fmt.Println("Perulangan pertama selesai")

	// Loop ke 2

	for i := 0 ; i < 7 ; i += 2 {
		fmt.Println("Perulangan ke ", i)
	}

	// Loop ke 3

	names := []string{"Timothy", "Dian", "Budi", "Joko"}

	for index, name := range names {
		fmt.Println("Index ke ", index, " = ", name)
	}

	person := map[string]string{
		"name" : "Timothy",
		"job" : "Software Engineer",
	}

	for key, value := range person {
		fmt.Println(key, " = ", value)
	}

	// Break and Continue
	for i := 0 ; i < 10 ; i++ {
		if i % 2 == 0 {
			continue
		}
		
		if i == 7 {
			break
		}	

		fmt.Println("Perulangan ke ", i)
	}
}