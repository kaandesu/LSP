# LSP

A language server protocol built with GO for educational purposes.

<small>
(
Check <a href="https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/">LSP specification</a> for more information.)</small>

---

Every progress will be tagged as checkpoints. Every tag will have a different README, more detailed explanation for each checkpoint. _To be updated_ contents of the checkpoints:

**[Checkpoint 1](./CHECKPOINT1.md)**

- [x] Basic `DecodeMessage` function
- [x] Basic `EncodeMessage` function
- [x] Basic `Split` function for the logger
- [x] Basic tests for the functions above
- [x] Basic `Logger` function
- [x] Starting `Stdin scanner` in main, with no-op `handleMessage` function

**Checkpoint 2**: <-- You are here<br>

- [x] Recieve basic messages from the lsp client and log them
- [x] Decoding the `initialize`
- [x] Initialize response
- [x] Text Document Synchronization

---

### Checkpoint 2

1- Basic logging when msg recieved from lsp client

```go
// main.go
func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}
```

#### Starting the LSP from Neovim

_in some nvim lua file_:

[Example](./appendix/lsp_first_message_cp2.lua) below, attaches the lsp client when buffer has the type 'markdown'. <br>
_`go build main.go` should be called, and used it's path for the `cmd`_

```lua
--- add this to somewhere with autocmds
local client = vim.lsp.start_client({
	name = "kaandesu/LSP",
	cmd = { "path/to/the/build/LSP/main" },
	filetypes = { "markdown" },
})
if not client then
	vim.notify("there is something wrong with the client")
end
vim.api.nvim_create_autocmd("FileType", {
	pattern = { "md", "markdown" },
	callback = function()
		if not client then
			vim.notify("there is something wrong with the client")
			return
		end
		vim.notify("custom lsp attached!")
		vim.lsp.buf_attach_client(0, client)
	end,
})

```

#### About Server Lifecycle - [specs](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#lifeCycleMessages)

1- **Initialize Request:** <br>

- The initialize request is sent as the first request from the client to the server.
- Server is not allowed to send any requests or notifications to the client until it has responded with an InitializeResult
- `initialize` request may only be sent once

- Afterwards, LSP shall respond with a `InitializeResult`

**NEXT:**

- [ ] TBD

---

**Honorable Mentions**

- [tjdevries/educationalsp](https://github.com/tjdevries/educationalsp)
- Brazil
