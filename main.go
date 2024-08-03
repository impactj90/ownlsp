package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/impactj90/ownlsp/rpc"
)

func main() {
    logger := getLogger("/Users/tolgaayvazoglu/projects/ownlsp/log.txt")
    logger.Println("Hey, i started")

	fmt.Println("hi")

	scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(msg)
	}
}

func handleMessage(_ any) {}

func getLogger(filename string) *log.Logger {
    logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
    if err != nil {
        panic("hey, you didnt give me a goold file. I panic just for educational reasons.")
    }

    return log.New(logfile, "[odnLsp]", log.Ldate|log.Ltime|log.Lshortfile)
}

