package main

import (
	"fmt"
	"os"

	"github.com/1stguardleft/discipline-gbatch/pkg/setting"
)

func init() {
	setting.Setup()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "Must specify a HTTP URL to get data from\n")
		os.Exit(1)
	}
}
