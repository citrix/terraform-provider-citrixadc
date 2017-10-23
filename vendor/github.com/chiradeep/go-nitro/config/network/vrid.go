package network

type Vrid struct {
	All               bool   `json:"all,omitempty"`
	Effectivepriority int    `json:"effectivepriority,omitempty"`
	Flags             int    `json:"flags,omitempty"`
	Id                int    `json:"id,omitempty"`
	Ifaces            string `json:"ifaces,omitempty"`
	Ipaddress         string `json:"ipaddress,omitempty"`
	Preemption        string `json:"preemption,omitempty"`
	Priority          int    `json:"priority,omitempty"`
	Sharing           string `json:"sharing,omitempty"`
	State             int    `json:"state,omitempty"`
	Tracking          string `json:"tracking,omitempty"`
	Type              string `json:"type,omitempty"`
}
