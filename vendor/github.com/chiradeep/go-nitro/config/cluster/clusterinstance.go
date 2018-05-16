package cluster

type Clusterinstance struct {
	Adminstate           string `json:"adminstate,omitempty"`
	Clid                 int    `json:"clid,omitempty"`
	Deadinterval         int    `json:"deadinterval,omitempty"`
	Hellointerval        int    `json:"hellointerval,omitempty"`
	Licensemismatch      bool   `json:"licensemismatch,omitempty"`
	Operationalpropstate string `json:"operationalpropstate,omitempty"`
	Operationalstate     string `json:"operationalstate,omitempty"`
	Preemption           string `json:"preemption,omitempty"`
	Propstate            string `json:"propstate,omitempty"`
	Rsskeymismatch       bool   `json:"rsskeymismatch,omitempty"`
	Status               string `json:"status,omitempty"`
}
