package main

import (
	"flag"
	"fmt"
	"os"
)

var test *int64

func init() {
	test = flag.Int64("test", 0, "command to review config template")
	flag.Parse()
}

func main(){
	gg := *test
	fmt.Println(gg)
	if gg < 4 {
		os.Exit(1)
	}
	if gg > 10 {
		os.Exit(0)
	}
}