package cs

type Cspolicy struct {
	Action       string `json:"action,omitempty"`
	Activepolicy bool   `json:"activepolicy,omitempty"`
	Bindhits     int    `json:"bindhits,omitempty"`
	Cspolicytype string `json:"cspolicytype,omitempty"`
	Domain       string `json:"domain,omitempty"`
	Hits         int    `json:"hits,omitempty"`
	Labelname    string `json:"labelname,omitempty"`
	Labeltype    string `json:"labeltype,omitempty"`
	Logaction    string `json:"logaction,omitempty"`
	Newname      string `json:"newname,omitempty"`
	Policyname   string `json:"policyname,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Rule         string `json:"rule,omitempty"`
	Url          string `json:"url,omitempty"`
	Vstype       int    `json:"vstype,omitempty"`
}
