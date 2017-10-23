package tm

type Tmtrafficpolicy struct {
	Action string `json:"action,omitempty"`
	Hits   int    `json:"hits,omitempty"`
	Name   string `json:"name,omitempty"`
	Rule   string `json:"rule,omitempty"`
}
