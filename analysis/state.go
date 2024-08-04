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

func (state *State) OpenDocument(document, text string) {
	state.Documents[document] = text
}
