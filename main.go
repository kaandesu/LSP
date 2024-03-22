package main

import (
	"LSP/rpc"
	"bufio"
	"log"
	"os"
)

func main() {
	logger := getLogger("/Users/kaanyilmaz/Desktop/Programming/gits/LSP/log.txt")
	logger.Println("Hey, I started!")

	scanner := bufio.NewScanner(os.Stdin)
	// scanner doesn't know how to read the LSP messages, so we split
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}
		// instead of massign the message, passing the method and the content
		handleMessage(logger, method, contents)
	}
}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("Recieved msg with method: %s", method)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(logfile, "[kaandesu/LSP]", log.Ldate|log.Ltime|log.Lshortfile)
}
