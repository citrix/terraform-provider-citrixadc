package network

type Nat64param struct {
	Nat64fragheader   string `json:"nat64fragheader,omitempty"`
	Nat64ignoretos    string `json:"nat64ignoretos,omitempty"`
	Nat64v6mtu        int    `json:"nat64v6mtu,omitempty"`
	Nat64zerochecksum string `json:"nat64zerochecksum,omitempty"`
	Td                int    `json:"td,omitempty"`
}
