package main

import (
	"os"
	"net/http"
	"fmt"
)



func main() {
	resp, err := http.Get("http://google.com/")
	
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	} else {
		fmt.Println(resp)
	}

}