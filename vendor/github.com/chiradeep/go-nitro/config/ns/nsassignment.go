package ns

type Nsassignment struct {
	Add            string `json:"Add,omitempty"`
	Append         string `json:"append,omitempty"`
	Clear          bool   `json:"clear,omitempty"`
	Comment        string `json:"comment,omitempty"`
	Hits           int    `json:"hits,omitempty"`
	Name           string `json:"name,omitempty"`
	Newname        string `json:"newname,omitempty"`
	Referencecount int    `json:"referencecount,omitempty"`
	Set            string `json:"set,omitempty"`
	Sub            string `json:"sub,omitempty"`
	Undefhits      int    `json:"undefhits,omitempty"`
	Variable       string `json:"variable,omitempty"`
}
