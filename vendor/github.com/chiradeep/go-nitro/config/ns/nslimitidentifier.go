package ns

type Nslimitidentifier struct {
	Computedtraptimeslice    int         `json:"computedtraptimeslice,omitempty"`
	Drop                     int         `json:"drop,omitempty"`
	Hits                     int         `json:"hits,omitempty"`
	Limitidentifier          string      `json:"limitidentifier,omitempty"`
	Limittype                string      `json:"limittype,omitempty"`
	Maxbandwidth             int         `json:"maxbandwidth,omitempty"`
	Mode                     string      `json:"mode,omitempty"`
	Ngname                   string      `json:"ngname,omitempty"`
	Referencecount           int         `json:"referencecount,omitempty"`
	Rule                     interface{} `json:"rule,omitempty"`
	Selectorname             string      `json:"selectorname,omitempty"`
	Threshold                int         `json:"threshold,omitempty"`
	Time                     int         `json:"time,omitempty"`
	Timeslice                int         `json:"timeslice,omitempty"`
	Total                    int         `json:"total,omitempty"`
	Trapscomputedintimeslice int         `json:"trapscomputedintimeslice,omitempty"`
	Trapsintimeslice         int         `json:"trapsintimeslice,omitempty"`
}
