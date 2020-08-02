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
		bs := make([]byte, 99999)

		n, err := resp.Body.Read(bs)
		fmt.Println(err)
		fmt.Println(n)
		fmt.Println(string(bs))
		
	}

}