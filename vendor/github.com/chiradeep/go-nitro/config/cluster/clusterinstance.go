package cluster

type Clusterinstance struct {
	Adminstate                 string `json:"adminstate,omitempty"`
	Backplanebasedview         string `json:"backplanebasedview,omitempty"`
	Clid                       int    `json:"clid,omitempty"`
	Clusternoheartbeatonnode   bool   `json:"clusternoheartbeatonnode,omitempty"`
	Clusternolinksetmbf        bool   `json:"clusternolinksetmbf,omitempty"`
	Clusternospottedip         bool   `json:"clusternospottedip,omitempty"`
	Deadinterval               int    `json:"deadinterval,omitempty"`
	Hellointerval              int    `json:"hellointerval,omitempty"`
	Heterogeneousflag          string `json:"heterogeneousflag,omitempty"`
	Inc                        string `json:"inc,omitempty"`
	Jumbonotsupported          bool   `json:"jumbonotsupported,omitempty"`
	Licensemismatch            bool   `json:"licensemismatch,omitempty"`
	Nodegroup                  string `json:"nodegroup,omitempty"`
	Nodegroupstatewarning      bool   `json:"nodegroupstatewarning,omitempty"`
	Nodepenummismatch          bool   `json:"nodepenummismatch,omitempty"`
	Operationalpropstate       string `json:"operationalpropstate,omitempty"`
	Operationalstate           string `json:"operationalstate,omitempty"`
	Penummismatch              bool   `json:"penummismatch,omitempty"`
	Preemption                 string `json:"preemption,omitempty"`
	Processlocal               string `json:"processlocal,omitempty"`
	Propstate                  string `json:"propstate,omitempty"`
	Quorumtype                 string `json:"quorumtype,omitempty"`
	Retainconnectionsoncluster string `json:"retainconnectionsoncluster,omitempty"`
	Rsskeymismatch             bool   `json:"rsskeymismatch,omitempty"`
	Status                     string `json:"status,omitempty"`
	Validmtu                   int    `json:"validmtu,omitempty"`
}
