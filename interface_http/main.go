package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com") //makes a request and assigns the object to the resp variable which is of type pointer response(struct)
	//Refer the Response struct image attached to this file. 'io.ReadCloser' is an interface and contains nested interfaces

	if err != nil {

		fmt.Println("Error:", err)

		os.Exit(1)

	}

	//	bs := make([]byte, 99999) //bs holds all the data that gets written into it by the read function. Make is a built in function that
	//initialises a byte slice with 99999 empty spaces

	// Read function is not set up to automatically resize the slice if the slice has already full. If we initialised 0 elements in bs
	// then Read function would regard it to already be full and thus we make an arbitrarily large byte slice for all the real source of data
	//to fit into

	//	resp.Body.Read(bs) //response body will take the byte slice bs and the html content and read it into byte slice

	//	fmt.Println(string(bs))

	// An alternative of the above

	fmt.Println("______________________ALTERNATIVE_______________________")

	lw := logWriter{}

	//	io.Copy(os.Stdout, resp.Body) // Prints the same thing as above
	// The second argument is the reader interface and the first argument is the writer interface

	io.Copy(lw, resp.Body) // Prints the same thing as above
}

//Custom type that implements the writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
