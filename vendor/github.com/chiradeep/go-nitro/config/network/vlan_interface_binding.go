package network

type Vlaninterfacebinding struct {
	Id         int    `json:"id,omitempty"`
	Ifnum      string `json:"ifnum,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Tagged     bool   `json:"tagged,omitempty"`
}
