package lsp

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
	// .. there's tons mor that goes here.. 
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
