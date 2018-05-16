package appfw

type Appfwpolicy struct {
	Comment     string `json:"comment,omitempty"`
	Hits        int    `json:"hits,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Name        string `json:"name,omitempty"`
	Newname     string `json:"newname,omitempty"`
	Policytype  string `json:"policytype,omitempty"`
	Profilename string `json:"profilename,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Undefhits   int    `json:"undefhits,omitempty"`
}
