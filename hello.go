package main

import (
	"encoding/json"
	"log"
	"net/http"

	"fmt"

	"rsc.io/quote"
)

type Greeting struct {
	Text     string `json:"text"`
	Language string `json:"language"`
	Code     string `json:"code"`
}

var greetings = []Greeting{
	Greeting{Text: "Hello World", Language: "English", Code: "en"},
	Greeting{Text: "你好 世界", Language: "Chinese", Code: "cn"},
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	enc.Encode(greetings)
}

func main() {
	fmt.Println(quote.Hello())
	// fmt.Println("👋, 🌎")
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
