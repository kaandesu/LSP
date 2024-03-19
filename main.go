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
		msg := scanner.Text()
		handleMessage(msg)
	}
}

func handleMessage(_ any) {}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(logfile, "[kaandesu/LSP]", log.Ldate|log.Ltime|log.Lshortfile)
}
