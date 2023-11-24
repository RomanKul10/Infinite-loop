package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"myapp/mylogger"
)

func main() {
	// reed input from the user  times and write it to a log

	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)

	go mylogger.LictenForLog(ch)

	fmt.Println("Enter something")

	for i := 0; i < 5; i++ {
		fmt.Print("->")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
	}
}
