package pcp

type Pcpserver struct {
	Ipaddress  string `json:"ipaddress,omitempty"`
	Name       string `json:"name,omitempty"`
	Pcpprofile string `json:"pcpprofile,omitempty"`
	Port       int    `json:"port,omitempty"`
}
