package network

type Interfacepair struct {
	Id     int         `json:"id,omitempty"`
	Ifaces string      `json:"ifaces,omitempty"`
	Ifnum  interface{} `json:"ifnum,omitempty"`
}
