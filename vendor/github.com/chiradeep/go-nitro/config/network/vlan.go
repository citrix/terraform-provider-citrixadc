package network

type Vlan struct {
	Aliasname          string `json:"aliasname,omitempty"`
	Id                 int    `json:"id,omitempty"`
	Ifaces             string `json:"ifaces,omitempty"`
	Ifnum              string `json:"ifnum,omitempty"`
	Ipv6dynamicrouting string `json:"ipv6dynamicrouting,omitempty"`
	Linklocalipv6addr  string `json:"linklocalipv6addr,omitempty"`
	Lsbitmap           int    `json:"lsbitmap,omitempty"`
	Lstagbitmap        int    `json:"lstagbitmap,omitempty"`
	Portbitmap         int    `json:"portbitmap,omitempty"`
	Rnat               bool   `json:"rnat,omitempty"`
	Sdxvlan            string `json:"sdxvlan,omitempty"`
	Tagbitmap          int    `json:"tagbitmap,omitempty"`
	Tagged             bool   `json:"tagged,omitempty"`
	Tagifaces          string `json:"tagifaces,omitempty"`
}
