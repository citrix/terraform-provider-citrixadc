package cmp

type Cmppolicy struct {
	Builtin            interface{} `json:"builtin,omitempty"`
	Clienttransactions int         `json:"clienttransactions,omitempty"`
	Clientttlb         int         `json:"clientttlb,omitempty"`
	Description        string      `json:"description,omitempty"`
	Expressiontype     string      `json:"expressiontype,omitempty"`
	Hits               int         `json:"hits,omitempty"`
	Isdefault          bool        `json:"isdefault,omitempty"`
	Name               string      `json:"name,omitempty"`
	Newname            string      `json:"newname,omitempty"`
	Reqaction          string      `json:"reqaction,omitempty"`
	Resaction          string      `json:"resaction,omitempty"`
	Rule               string      `json:"rule,omitempty"`
	Rxbytes            int         `json:"rxbytes,omitempty"`
	Servertransactions int         `json:"servertransactions,omitempty"`
	Serverttlb         int         `json:"serverttlb,omitempty"`
	Txbytes            int         `json:"txbytes,omitempty"`
}
