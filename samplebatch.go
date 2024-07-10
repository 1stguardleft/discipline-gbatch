package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetchRemoteResource(url string) ([]byte, error) {
	//reqBody := bytes.NewBufferString("")
	//r, err := http.Post(urlocator, "multipart/form-data", reqBody)
	/*
			  mktId = 'STK',
		  trdDd = '20210108',
		  money = '1',
		  csvxls_isNo = 'false',
		  name = 'fileDown',
		  url = 'dbms/MDC/STAT/standard/MDCSTAT03901'
	*/
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	return io.ReadAll(r.Body)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "Must specify a HTTP URL to get data from\n")
		os.Exit(1)
	}

	body, err := fetchRemoteResource(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s\n", body)

}
