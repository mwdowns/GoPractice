package main

import (
	"fmt"
	"io"
	"os"
)

type fileReader struct{}

func main() {

	resp, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fr := fileReader{}
	io.Copy(fr, resp)
}

func (fileReader) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("the byte slice was this long: ", len(bs))
	return len(bs), nil
}
