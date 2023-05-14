package main

import (
	"fmt"
	"os"
)

func main() {

	var s []string = os.Args
	for i, v := range s {

		fmt.Println(i, v)
	}

}
