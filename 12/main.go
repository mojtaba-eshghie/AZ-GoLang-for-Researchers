package main

import (
	"io"
	"os"
	"net/http"
	"fmt"
)

type logWriter struct {

}

func (logWriter) Write(bs []byte) (int, error){
	//just by defining the logWriter and associating logWriter to this Write
	// function, logWriter is now implementing the Writer interface.
	fmt.Println(string(bs))


	// to implement the return we have to take a look at the Write interface documentation 
	// to implement it according to the docs
	// here, we should return (number of bytes written, error)
	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://google.com/")
	
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	} 

	io.Copy(os.Stdout, resp.Body)

	////////////////////////////// second part //////////////////////////

	//let's instantiate this struct now
	lw := logWriter{}
	
	//let's do this:
	io.Copy(lw, resp.Body)

}

