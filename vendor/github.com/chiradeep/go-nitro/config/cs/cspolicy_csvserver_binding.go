package cs

type Cspolicycsvserverbinding struct {
	Action       string `json:"action,omitempty"`
	Domain       string `json:"domain,omitempty"`
	Hits         int    `json:"hits,omitempty"`
	Labelname    string `json:"labelname,omitempty"`
	Labeltype    string `json:"labeltype,omitempty"`
	Pihits       int    `json:"pihits,omitempty"`
	Pipolicyhits int    `json:"pipolicyhits,omitempty"`
	Policyname   string `json:"policyname,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Url          string `json:"url,omitempty"`
}
