package network

type Inatparam struct {
	Nat46fragheader   string `json:"nat46fragheader,omitempty"`
	Nat46ignoretos    string `json:"nat46ignoretos,omitempty"`
	Nat46v6mtu        int    `json:"nat46v6mtu,omitempty"`
	Nat46v6prefix     string `json:"nat46v6prefix,omitempty"`
	Nat46zerochecksum string `json:"nat46zerochecksum,omitempty"`
	Td                int    `json:"td,omitempty"`
}
