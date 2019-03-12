package rdp

type Rdpserverprofile struct {
	Builtin        interface{} `json:"builtin,omitempty"`
	Name           string      `json:"name,omitempty"`
	Psk            string      `json:"psk,omitempty"`
	Rdpip          string      `json:"rdpip,omitempty"`
	Rdpport        int         `json:"rdpport,omitempty"`
	Rdpredirection string      `json:"rdpredirection,omitempty"`
}
