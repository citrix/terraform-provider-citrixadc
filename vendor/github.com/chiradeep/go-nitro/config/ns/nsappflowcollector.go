package ns

type Nsappflowcollector struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Name      string `json:"name,omitempty"`
	Port      int    `json:"port,omitempty"`
}
