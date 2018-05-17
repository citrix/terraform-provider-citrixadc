package co

type Copolicy struct {
	Action bool   `json:"action,omitempty"`
	Name   string `json:"name,omitempty"`
	Rule   bool   `json:"rule,omitempty"`
}
