package main

import (
	"GoLangPractice/src/chans"
	"GoLangPractice/src/maps"
	"fmt"
	"os"
)

func main() {

	isLoopMode := os.Args[1] == "true"
	loopMode := ""
	if isLoopMode {
		loopMode = "loop"
	} else {
		loopMode = "no loop"
	}

	fmt.Println("function testChanCap() will run in " + loopMode + " mode")
	testChanCap(isLoopMode)

	testHashMap()
}

func testChanCap(isLoopMode bool) {
	ch := make(chan string, 4)
	resultReport := ""

	if isLoopMode {
		for i := 0; i < 10; i++ {
			go chans.Run(i, ch)
		}
	} else {
		go chans.Run(0, ch)
		go chans.Run(1, ch)
		go chans.Run(2, ch)
		go chans.Run(3, ch)
		go chans.Run(4, ch)
		go chans.Run(5, ch)
		go chans.Run(6, ch)
		go chans.Run(7, ch)
		go chans.Run(8, ch)
		go chans.Run(9, ch)
	}

	for i := 0; i < 10; i++ {
		resultReport += <-ch
	}
	fmt.Println(resultReport)
}

func testHashMap() {
	ch := make(chan string, 4)
	go testMapWithChan("firstValue", "secondValue", ch)
	go testMapWithChan("thirdValue", "forthValue", ch)
	fmt.Println(<-ch + <-ch)
}

func testMapWithChan(value1 string, value2 string, ch chan string) {
	testMap := new(maps.HashMap)
	testMap.Put("1", value1)
	testMap.Put("2", value2)
	ch <- "The first value of the map in the ch is: " + testMap.Get("1") + "\n" + "the second value of the map in the ch is: " + testMap.Get("2") + "\n"
}
