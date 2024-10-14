package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Random_Quote struct {
	Q string `json : "a"`
	A string `json : "q"`
}

func main() {

	response, err := http.Get("https://zenquotes.io/api/random")
	// response = "1";
	// fmt.Printf("Response is %s\n", response.Body)
	if err != nil {
		fmt.Printf("Error is %d\n", err)
	}

	body, err := io.ReadAll(response.Body)
	// fmt.Printf("%s\n", string(body))
	// fmt.Printf("%s", map(string(body[0])))
	// fmt.Printf("%s", body[0]["q"])
	var quote []Random_Quote
	err = json.Unmarshal(body, &quote)
	fmt.Printf("Quote : %s\nAuthor : %s", quote[0].Q, quote[0].A)

}
