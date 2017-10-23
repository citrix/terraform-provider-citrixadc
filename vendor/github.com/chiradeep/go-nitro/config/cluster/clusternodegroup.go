package cluster

type Clusternodegroup struct {
	Activelist               interface{} `json:"activelist,omitempty"`
	Backuplist               interface{} `json:"backuplist,omitempty"`
	Backupnodemask           int         `json:"backupnodemask,omitempty"`
	Boundedentitiescntfrompe int         `json:"boundedentitiescntfrompe,omitempty"`
	Currentnodemask          int         `json:"currentnodemask,omitempty"`
	Name                     string      `json:"name,omitempty"`
	Strict                   string      `json:"strict,omitempty"`
}
