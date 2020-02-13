package ns

type Nslicenseserver struct {
	Forceupdateip   bool   `json:"forceupdateip,omitempty"`
	Gptimeleft      int    `json:"gptimeleft,omitempty"`
	Grace           int    `json:"grace,omitempty"`
	Licenseserverip string `json:"licenseserverip,omitempty"`
	Nodeid          int    `json:"nodeid,omitempty"`
	Port            int    `json:"port,omitempty"`
	Servername      string `json:"servername,omitempty"`
	Status          int    `json:"status,omitempty"`
}
