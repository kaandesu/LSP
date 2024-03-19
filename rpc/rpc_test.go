package rpc_test

import (
	"LSP/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	/* 16 is the number of characters coming after that number
	/* which is the content as defined in rpc.go */
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})
	if expected != actual {
		t.Fatalf("Ecpected: %s, Actual: %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	// checking the "method"
	incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
	contentLength := len(content)

	if err != nil {
		t.Fatal(err)
	}
	if contentLength != 15 {
		t.Fatalf("Expected: 16, Got: %d", contentLength)
	}

	if method != "hi" {
		t.Fatalf("Expected: 'hi', Got: %s ", method)
	}
}
