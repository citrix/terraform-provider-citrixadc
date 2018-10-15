package network

type Vrid struct {
	All                  bool   `json:"all,omitempty"`
	Effectivepriority    int    `json:"effectivepriority,omitempty"`
	Flags                int    `json:"flags,omitempty"`
	Id                   int    `json:"id,omitempty"`
	Ifaces               string `json:"ifaces,omitempty"`
	Ipaddress            string `json:"ipaddress,omitempty"`
	Operationalownernode int    `json:"operationalownernode,omitempty"`
	Ownernode            int    `json:"ownernode,omitempty"`
	Preemption           string `json:"preemption,omitempty"`
	Preemptiondelaytimer int    `json:"preemptiondelaytimer,omitempty"`
	Priority             int    `json:"priority,omitempty"`
	Sharing              string `json:"sharing,omitempty"`
	State                int    `json:"state,omitempty"`
	Trackifnumpriority   int    `json:"trackifnumpriority,omitempty"`
	Tracking             string `json:"tracking,omitempty"`
	Type                 string `json:"type,omitempty"`
}
