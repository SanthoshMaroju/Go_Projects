package main

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
)

func main() {
	fmt.Printf("Hi, Learning Go")
	response, err := http.Get("https://zenquotes.io/api/random")
	// response = "1";
	fmt.Printf("Response is %s\n", response.Body)
	fmt.Printf("Error is %d\n", err)

	text, err := io.ReadAll(response.Body)
	// fmt.Printf("%s\n", text[0]['q'])
	fmt.Printf("%s", reflect.TypeOf(text))
}
