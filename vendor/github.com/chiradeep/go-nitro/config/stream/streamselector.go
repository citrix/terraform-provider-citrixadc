package stream

type Streamselector struct {
	Name string      `json:"name,omitempty"`
	Rule interface{} `json:"rule,omitempty"`
}
