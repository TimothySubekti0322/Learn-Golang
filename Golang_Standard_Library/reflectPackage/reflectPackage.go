package reflectpackage

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"` // tag
}

type Person struct {
	Name string `required:"true"`
	Age  int   `required:"true"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("Type :", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name, "with type", valueField.Type)
		fmt.Println("Tag :", valueField.Tag)
		fmt.Println("Required :", valueField.Tag.Get("required"))
	}
}

func MainReflectPackage() {
	readField(Sample{"Timothy"})
	readField(Person{"John Doe", 30})

	person := Person{"Timothy", 30}
	fmt.Println("isValid Person :", IsValid(person))

}