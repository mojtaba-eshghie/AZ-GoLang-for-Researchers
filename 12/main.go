package main

import (
	"io"
	"os"
	"net/http"
	"fmt"
)



func main() {
	resp, err := http.Get("http://google.com/")
	
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	} 

	io.Copy(os.Stdout, resp.Body)

}