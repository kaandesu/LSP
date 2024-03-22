# LSP

A language server protocol built with GO for educational purposes.

<small>
(
Check <a href="https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/">LSP specification</a> for more information.)</small>

---

Every progress will be tagged as checkpoints. Every tag will have a different README, more detailed explanation for each checkpoint. _To be updated_ contents of the checkpoints:

**Checkpoint 1** <-- You are here:

- [x] Basic `DecodeMessage` function
- [x] Basic `EncodeMessage` function
- [x] Basic `Split` function for the logger
- [x] Basic tests for the functions above
- [x] Basic `Logger` function
- [x] Starting `Stdin scanner` in main, with no-op `handleMessage` function

**Checkpoint 2**: <br>
[*TBD*]

---

### Checkpoint 1

---

1- Creating RPC package

```bash
go mod init [NAME]
mkdir rpc
touch rpc.go
```

#### rpc.go content:

- `func EncodedMessage(msg any) string`

2- Creating `rpc_test.go`

```go
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
```

3- `rpc.go` - Creating `func DecodeMessage(msg []byte)`

- cutting the bytes to get data from the content, we cut by \r \n \r \n see [LSP specification](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/)

```
Content-Length: ...\r\n
\r\n
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "textDocument/completion",
  "params": {
    ...
  }
}
```

```go
func DecodeMessage(msg []byte) (int, error) {
	/*
	 * bytes.Cut: it takes some slice, takes some seperator
	 * gives all the bytes before that and after that, and if we found it
	 */
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, errors.New("Did not find seperator")
	}
	// we found "Content-Length: <number>..."
	contentLengthBytes := header[len("Content-Length: "):]
  // converting from bytes, i think
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, err
	}

	// TODO: i'll get to this later
	_ = content
	return contentLength, nil
}
```

4- Add decoding test to RPC_Test.go

```go
func TestDecode(t *testing.T) {
	incomingMessage := "Content-Length: 16\r\n\r\n{\"Testing\":true}"

	contentLength, err := rpc.DecodeMessage([]byte(incomingMessage))
	if err != nil {
		t.Fatal(err)
	}
	if contentLength != 16 {
		t.Fatalf("Expected: 16, Got: %d", contentLength)
	}
}
```

#### Decoding Method

example, we want to return this value:

```json
...
"method": "textDocument/completion",
...
```

in **rpc.go**:

- `DecodeMessage` now returns the decoded method as first argument

```go

type BaseMessage struct {
	Method string `json:"method"`
}

func DecodeMessage(msg []byte) (string, int, error){
  ...

 	var baseMessage BaseMessage

	if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return "", 0, err
	}
	return baseMessage.Method, contentLength, nil
```

**SOME CHANGES**:

- `DecodeMessage` returns: string (message of method), []byte (which is the actual content), error

```go
func DecodeMessage(msg []byte) (string, []byte, error) {
  ...

	var baseMessage BaseMessage

	if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return "", nil, err
	}
	return baseMessage.Method, content[:contentLength], nil
}
```

#### Building and Basic Logging

**main.go**:

1- Adding the `scanner`

```go
func main(){
  // when scanner.Scan() called, it will wait for a standard input
  scanner := bufio.NewScanner(os.Stdin)


  // scanner doesn't know how to read the LSP messages, so we pass a split function
  scanner.Split(rpc.Split)


  // infinite while loop, that waits for a stdin to invoke
  for scanner.Scan(){
    msg : scanner.Text()

    // not-yet implemented function
    handleMessage(msg)
  }


  func handleMessage(_ any) {}

}
```

in `rpc.go`:

```go
// for reference: type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil
	}
	// we found "Content-Length: <number>..."
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, err
	}
	if len(content) < contentLength {
		return 0, nil, nil
	}
	totalLength := len(header) + 4 + contentLength
	return totalLength, data[:totalLength], nil
}
```

2- Adding the `Logger`

```go
func main(){
  logger := getLogger("path/to/the/log/file.txt")

  // writing something to that file
  logger.Println("Hey, logger is started")

  .....

}

func getLogger(filename string) *log.Logger{
  // creates every time and has write only permission
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(logfile, "[kaandesu/LSP]", log.Ldate|log.Ltime|log.Lshortfile)
}

```

3- Starting the logger to test

```bash
go run main.go
```

4- Check the specified log file if it worked

---

**NEXT:**

- [ ] Add actual basic logging to the message handler

**Honorable Mentions**

- [tjdevries/educationalsp](https://github.com/tjdevries/educationalsp)
- Brazil
