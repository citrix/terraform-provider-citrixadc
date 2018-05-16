package network

type Vlaninterfacebinding struct {
	Id     int    `json:"id,omitempty"`
	Ifnum  string `json:"ifnum,omitempty"`
	Tagged bool   `json:"tagged,omitempty"`
}
