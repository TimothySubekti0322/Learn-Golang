package containerListPackage

import (
	"container/list"
	"fmt"
)

func MainContainerListPackage() {
	var data *list.List = list.New()

	data.PushBack("A") // A
	data.PushFront("B") // B A
	data.PushBack("C") // B A C
	data.PushFront("D") // D B A C

	// Let's search for the element "B"

	found := false
	e := data.Front()

	for !found {
		if(e.Value == "B") {
			data.InsertAfter("X", e) // D B X A C
			found = true
		} else {
			e = e.Next()
		}
	}

	for e := data.Front(); e != nil; e = e.Next() {
		var printedValue string = e.Value.(string) + " "
		fmt.Print(printedValue)
	}

	fmt.Print("\n")

	// kalau ujung elemen di Next akan nil
}