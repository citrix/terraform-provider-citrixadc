package dns

type Dnspolicy64 struct {
	Action      string `json:"action,omitempty"`
	Description string `json:"description,omitempty"`
	Hits        int    `json:"hits,omitempty"`
	Labelname   string `json:"labelname,omitempty"`
	Labeltype   string `json:"labeltype,omitempty"`
	Name        string `json:"name,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Undefhits   int    `json:"undefhits,omitempty"`
}
