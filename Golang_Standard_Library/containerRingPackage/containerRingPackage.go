package containerRingPackage

import (
	"container/ring"
	"fmt"
)

func MainContainerRingPackage() {
	data := ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = string(i + 65)
		data = data.Next()
	}

	var firstData *ring.Ring = data

	var done bool = false

	e := data

	for !done {
		fmt.Println(e)
		if e.Next() == firstData {
			done = true
		} else {
			e = e.Next()
		}
	}
}