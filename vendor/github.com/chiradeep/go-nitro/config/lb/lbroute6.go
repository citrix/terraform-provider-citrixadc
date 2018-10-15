package lb

type Lbroute6 struct {
	Flags       string `json:"flags,omitempty"`
	Gatewayname string `json:"gatewayname,omitempty"`
	Network     string `json:"network,omitempty"`
	Td          int    `json:"td,omitempty"`
}
