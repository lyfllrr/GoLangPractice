package chans

import (
	"fmt"
	"strconv"
)

func Run(i int, ch chan string) {
	fmt.Println("There're " + strconv.Itoa(len(ch)) + " channels awaiting.")
	//time.Sleep(10 * time.Millisecond)
	ch <- "ch" + strconv.Itoa(i) + " is finished. \n"
}
