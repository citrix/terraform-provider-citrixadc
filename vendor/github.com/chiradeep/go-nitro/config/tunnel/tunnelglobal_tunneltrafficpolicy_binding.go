package tunnel

type Tunnelglobaltunneltrafficpolicybinding struct {
	Builtin    interface{} `json:"builtin,omitempty"`
	Policyname string      `json:"policyname,omitempty"`
	Priority   int         `json:"priority,omitempty"`
	State      string      `json:"state,omitempty"`
}
