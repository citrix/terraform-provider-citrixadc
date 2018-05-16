package filter

type Filterpolicy struct {
	Hits      int    `json:"hits,omitempty"`
	Name      string `json:"name,omitempty"`
	Reqaction string `json:"reqaction,omitempty"`
	Resaction string `json:"resaction,omitempty"`
	Rule      string `json:"rule,omitempty"`
}
