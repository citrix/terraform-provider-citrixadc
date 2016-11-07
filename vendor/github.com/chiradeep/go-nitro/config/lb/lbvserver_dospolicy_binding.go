package lb

type Lbvserverdospolicybinding struct {
	Name       string `json:"name,omitempty"`
	Policyname string `json:"policyname,omitempty"`
	Priority   int    `json:"priority,omitempty"`
}
