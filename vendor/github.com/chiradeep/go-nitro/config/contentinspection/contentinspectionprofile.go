package contentinspection

type Contentinspectionprofile struct {
	Egressinterface  string      `json:"egressinterface,omitempty"`
	Egressvlan       int         `json:"egressvlan,omitempty"`
	Ingressinterface string      `json:"ingressinterface,omitempty"`
	Ingressvlan      int         `json:"ingressvlan,omitempty"`
	Mode             interface{} `json:"mode,omitempty"`
	Name             string      `json:"name,omitempty"`
	Type             string      `json:"type,omitempty"`
}
