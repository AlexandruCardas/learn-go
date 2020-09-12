package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	message := "Hello There"
	// Can use postfix after using a dot and CTRL + SPACE twice in order to get the methods that can be used for this main
	strings.ToTitle(message)

	// type `if` and then chose an option

	// While inside the brackets calling a function, press CTRL + P to get info on the parameters
	myFunc(3, "hi")

	// Can also be used for types
	_ = demo{
		Field1:    0,
		Field2:    "",
		SubStruct: SubStruct{},
	}

	_ = demo{0, "", SubStruct{SubField1: 3}}

	// Inject language with ALT + ENTER, close the editor with CTRL + F4
	json := `{"j" : "whatever"}`
	fmt.Println(json)

	//language=GoTemplate
	message1 := `{{- /*gotype: github.com/dlsniper/tipsandtricks/editing/tip004.Output*/ -}}
各位 | 晚上好!
We have | Gophers here today!
`

	printMessage(message1, "Gopher")
}

type Output struct {
	Message string
	Count   int
}

func printMessage(msg, name string) {
	t := template.Must(template.New("output").Parse(msg))

	o := Output{Message: name}

	err := t.Execute(os.Stdout, o)
	if err != nil {
		log.Fatalln(err)
	}
}

type (
	SubStruct struct {
		SubField1 int
	}

	demo struct {
		Field1 int
		Field2 string
		SubStruct
	}
)

func _() (int, error) {
	fmt.Print("lol")
	_ = io.ReadWriteCloser(nil) // Only need to type rewrcl for the entire function to autocomplete

	// just type ret and everything else almost gets completed
	return 0, errors.New("error")
}

func convertToString(num int) string {
	return fmt.Sprintf("%d", num) // can use cyclical typing to find the name of convertToString (type first word and press ALT + /)
}

func myFunc(n int, s string) {

}
