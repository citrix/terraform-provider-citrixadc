package tunnel

type Tunneltrafficpolicy struct {
	Action             string      `json:"action,omitempty"`
	Builtin            interface{} `json:"builtin,omitempty"`
	Clienttransactions int         `json:"clienttransactions,omitempty"`
	Clientttlb         int         `json:"clientttlb,omitempty"`
	Comment            string      `json:"comment,omitempty"`
	Expressiontype     string      `json:"expressiontype,omitempty"`
	Hits               int         `json:"hits,omitempty"`
	Isdefault          bool        `json:"isdefault,omitempty"`
	Logaction          string      `json:"logaction,omitempty"`
	Name               string      `json:"name,omitempty"`
	Newname            string      `json:"newname,omitempty"`
	Rule               string      `json:"rule,omitempty"`
	Rxbytes            int         `json:"rxbytes,omitempty"`
	Servertransactions int         `json:"servertransactions,omitempty"`
	Serverttlb         int         `json:"serverttlb,omitempty"`
	Txbytes            int         `json:"txbytes,omitempty"`
	Undefhits          int         `json:"undefhits,omitempty"`
}
