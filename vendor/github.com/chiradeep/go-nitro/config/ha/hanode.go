package ha

type Hanode struct {
	Completedfliptime int    `json:"completedfliptime,omitempty"`
	Curflips          int    `json:"curflips,omitempty"`
	Deadinterval      int    `json:"deadinterval,omitempty"`
	Disifaces         string `json:"disifaces,omitempty"`
	Enaifaces         string `json:"enaifaces,omitempty"`
	Failsafe          string `json:"failsafe,omitempty"`
	Flags             int    `json:"flags,omitempty"`
	Haheartbeatifaces string `json:"haheartbeatifaces,omitempty"`
	Hamonifaces       string `json:"hamonifaces,omitempty"`
	Haprop            string `json:"haprop,omitempty"`
	Hastatus          string `json:"hastatus,omitempty"`
	Hasync            string `json:"hasync,omitempty"`
	Hellointerval     int    `json:"hellointerval,omitempty"`
	Id                int    `json:"id,omitempty"`
	Ifaces            string `json:"ifaces,omitempty"`
	Inc               string `json:"inc,omitempty"`
	Ipaddress         string `json:"ipaddress,omitempty"`
	Masterstatetime   int    `json:"masterstatetime,omitempty"`
	Maxflips          int    `json:"maxflips,omitempty"`
	Maxfliptime       int    `json:"maxfliptime,omitempty"`
	Name              string `json:"name,omitempty"`
	Netmask           string `json:"netmask,omitempty"`
	Pfifaces          string `json:"pfifaces,omitempty"`
	Routemonitor      string `json:"routemonitor,omitempty"`
	Routemonitorstate string `json:"routemonitorstate,omitempty"`
	Ssl2              string `json:"ssl2,omitempty"`
	State             string `json:"state,omitempty"`
	Syncvlan          int    `json:"syncvlan,omitempty"`
}
