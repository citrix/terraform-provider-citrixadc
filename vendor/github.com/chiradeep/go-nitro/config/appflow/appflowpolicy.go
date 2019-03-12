package appflow

type Appflowpolicy struct {
	Action      string `json:"action,omitempty"`
	Comment     string `json:"comment,omitempty"`
	Description string `json:"description,omitempty"`
	Hits        int    `json:"hits,omitempty"`
	Name        string `json:"name,omitempty"`
	Newname     string `json:"newname,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
	Undefhits   int    `json:"undefhits,omitempty"`
}
