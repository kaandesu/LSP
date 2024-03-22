package lsp

type Request struct {
	RPC    string `json:"jsonprc"`
	Method string `json:"method"`
	ID     int    `json:"id"`

	// WARN: will just specify the type of the params in all the Request types
	// Params ...
}

type Response struct {
	ID  *int   `json:"id,omitempty"`
	RPC string `json:"jsonprc"`

	// WARN: this will also have 'Result' and 'Error'
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}
