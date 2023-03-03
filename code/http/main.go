package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func (w logWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	fmt.Println("bytes", len(p))
	panic("implement me")
	return len(p), nil
}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	//bs := make([]byte, 999)
	//resp.Body.Read(bs)
	//
	//fmt.Println(string(bs))

	// io.Copy(Writer, Reader)
	// http.Responseはreaderをinterfaceしている

	lw := logWriter{}
	//io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)
}
