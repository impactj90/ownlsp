package analysis

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

