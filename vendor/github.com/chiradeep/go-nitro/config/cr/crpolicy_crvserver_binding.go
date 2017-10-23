package cr

type Crpolicycrvserverbinding struct {
	Domain                 string `json:"domain,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Pihits                 int    `json:"pihits,omitempty"`
	Pipolicyhits           int    `json:"pipolicyhits,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}
