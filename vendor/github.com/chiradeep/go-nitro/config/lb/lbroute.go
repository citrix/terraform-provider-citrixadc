package lb

type Lbroute struct {
	Flags       string `json:"flags,omitempty"`
	Gatewayname string `json:"gatewayname,omitempty"`
	Netmask     string `json:"netmask,omitempty"`
	Network     string `json:"network,omitempty"`
	Td          int    `json:"td,omitempty"`
}
