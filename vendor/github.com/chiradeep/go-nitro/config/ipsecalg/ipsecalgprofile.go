package ipsecalg

type Ipsecalgprofile struct {
	Connfailover      string `json:"connfailover,omitempty"`
	Espgatetimeout    int    `json:"espgatetimeout,omitempty"`
	Espsessiontimeout int    `json:"espsessiontimeout,omitempty"`
	Ikesessiontimeout int    `json:"ikesessiontimeout,omitempty"`
	Name              string `json:"name,omitempty"`
}
