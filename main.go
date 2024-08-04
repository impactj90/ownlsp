package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/impactj90/ownlsp/analysis"
	"github.com/impactj90/ownlsp/lsp"
	"github.com/impactj90/ownlsp/rpc"
)

func main() {
	logger := getLogger("/Users/tolgaayvazoglu/projects/ownlsp/log.txt")
	logger.Println("Hey, i started")

	fmt.Println("hi")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got en error: %s", err)
			continue
		}

		handleMessage(logger, state, method, contents)
	}
}

func handleMessage(logger *log.Logger, state analysis.State, method string, contents []byte) {
	logger.Printf("Received msg with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("We could not parse this: %s", err)
		}
		logger.Printf("Connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

		// hey.. let's reply
		msg := lsp.NewInitializedResponse(request.ID)
		reply := rpc.EncodeMessage(msg)

		writer := os.Stdout
		writer.Write([]byte(reply))

		logger.Print("Sent the reply")

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("We could not parse this: %s", err)
		}
		logger.Printf("Opened: %s %s",
			request.Params.TextDocument.URI,
			request.Params.TextDocument.Text,
		)

		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, you didnt give me a goold file. I panic just for educational reasons.")
	}

	return log.New(logfile, "[ownLsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
