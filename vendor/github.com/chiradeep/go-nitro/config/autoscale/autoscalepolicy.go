package autoscale

type Autoscalepolicy struct {
	Action    string `json:"action,omitempty"`
	Comment   string `json:"comment,omitempty"`
	Hits      int    `json:"hits,omitempty"`
	Logaction string `json:"logaction,omitempty"`
	Name      string `json:"name,omitempty"`
	Newname   string `json:"newname,omitempty"`
	Priority  int    `json:"priority,omitempty"`
	Rule      string `json:"rule,omitempty"`
	Undefhits int    `json:"undefhits,omitempty"`
}
