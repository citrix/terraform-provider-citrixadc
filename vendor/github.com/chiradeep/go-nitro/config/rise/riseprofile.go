package rise

type Riseprofile struct {
	Deviceid      string `json:"deviceid,omitempty"`
	Ifnum         string `json:"ifnum,omitempty"`
	Ipaddress     string `json:"ipaddress,omitempty"`
	Issu          string `json:"issu,omitempty"`
	Mode          string `json:"mode,omitempty"`
	Profilename   string `json:"profilename,omitempty"`
	Servicename   string `json:"servicename,omitempty"`
	Slotid        int    `json:"slotid,omitempty"`
	Slotno        int    `json:"slotno,omitempty"`
	Status        string `json:"status,omitempty"`
	Vdcid         int    `json:"vdcid,omitempty"`
	Vlan          int    `json:"vlan,omitempty"`
	Vlancfgstatus string `json:"vlancfgstatus,omitempty"`
	Vlangroupid   int    `json:"vlangroupid,omitempty"`
	Vpcid         int    `json:"vpcid,omitempty"`
}
