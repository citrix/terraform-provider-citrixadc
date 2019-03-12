package lb

type Lbvservercsvserverbinding struct {
	Cachetype     string `json:"cachetype,omitempty"`
	Cachevserver  string `json:"cachevserver,omitempty"`
	Hits          int    `json:"hits,omitempty"`
	Name          string `json:"name,omitempty"`
	Pipolicyhits  int    `json:"pipolicyhits,omitempty"`
	Policyname    string `json:"policyname,omitempty"`
	Policysubtype int    `json:"policysubtype,omitempty"`
	Priority      int    `json:"priority,omitempty"`
}
