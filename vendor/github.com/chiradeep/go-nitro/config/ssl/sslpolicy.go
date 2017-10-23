package ssl

type Sslpolicy struct {
	Action      string `json:"action,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Description string `json:"description,omitempty"`
	Hits        int    `json:"hits,omitempty"`
	Name        string `json:"name,omitempty"`
	Policytype  string `json:"policytype,omitempty"`
	Reqaction   string `json:"reqaction,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
	Undefhits   int    `json:"undefhits,omitempty"`
}
