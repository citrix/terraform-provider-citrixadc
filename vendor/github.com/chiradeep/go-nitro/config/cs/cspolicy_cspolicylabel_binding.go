package cs

type Cspolicycspolicylabelbinding struct {
	Domain     string `json:"domain,omitempty"`
	Hits       int    `json:"hits,omitempty"`
	Labelname  string `json:"labelname,omitempty"`
	Labeltype  string `json:"labeltype,omitempty"`
	Policyname string `json:"policyname,omitempty"`
	Priority   int    `json:"priority,omitempty"`
	Url        string `json:"url,omitempty"`
}
