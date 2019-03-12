package appflow

type Appflowcollector struct {
	Ipaddress  string `json:"ipaddress,omitempty"`
	Name       string `json:"name,omitempty"`
	Netprofile string `json:"netprofile,omitempty"`
	Newname    string `json:"newname,omitempty"`
	Port       int    `json:"port,omitempty"`
	State      string `json:"state,omitempty"`
	Transport  string `json:"transport,omitempty"`
}
