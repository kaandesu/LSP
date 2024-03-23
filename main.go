package main

import (
	"LSP/lsp"
	"LSP/rpc"
	"bufio"
	"encoding/json"
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

	switch method {
	case "initialize":
		// As seen in the lsp spesification, after getting the first 'initialize' message we need to respond back
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Hey, we couldn't parse this: %s", err)
		}
		logger.Printf("Connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

		// let's reply, will change later
		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)
		// fmt.Printf("repllllyy : %s", reply)

		writer := os.Stdout
		writer.Write([]byte(reply))
		logger.Print("Sent the reply")

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Hey, we couldn't parse this: %s", err)
		}
		logger.Printf("Opened: %s, %s", request.Params.TextDocument.URI, request.Params.TextDocument.Text)

	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(logfile, "[kaandesu/LSP]", log.Ldate|log.Ltime|log.Lshortfile)
}
