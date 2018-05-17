package network

type Vlanlinksetbinding struct {
	Id     int    `json:"id,omitempty"`
	Ifnum  string `json:"ifnum,omitempty"`
	Tagged bool   `json:"tagged,omitempty"`
}
