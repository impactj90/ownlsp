package analysis

import (
	"fmt"

	"github.com/impactj90/ownlsp/lsp"
)

type State struct {
	// map of filenames to contents
	Documents map[string]string
}

func NewState() State {
	return State{
		Documents: map[string]string{},
	}
}

func (state *State) OpenDocument(uri, text string) {
	state.Documents[uri] = text
}

func (state *State) UpdateDocument(uri, text string) {
	state.Documents[uri] = text
}

func (state *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	// in real life this would look up the type in our type analysis code..

	document := state.Documents[uri]

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File :%s, Characters: %d", uri, len(document)),
		},
	}
}

func (state *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
			},
		},
	}
}
