package main

import (
	"fmt"
	"net/http"

	_ "github.com/cohunter/preresolve"
)

func main() {
	resp, err := http.Get("https://github.com/cohunter")
	fmt.Printf("%+v, %+v", err, resp)
}
