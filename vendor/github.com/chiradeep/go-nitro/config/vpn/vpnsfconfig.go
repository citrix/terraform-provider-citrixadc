package vpn

type Vpnsfconfig struct {
	Filename string      `json:"filename,omitempty"`
	Vserver  interface{} `json:"vserver,omitempty"`
}
