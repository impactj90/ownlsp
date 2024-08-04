package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`

	// we will just specifigy the type of the params in all the rquest types later
	// Params ...
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id,ommitempty"`

	// Result
	// Error
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}
