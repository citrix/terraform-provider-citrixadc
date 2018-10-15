package lsn

type Lsntransportprofile struct {
	Finrsttimeout        int    `json:"finrsttimeout,omitempty"`
	Groupsessionlimit    int    `json:"groupsessionlimit,omitempty"`
	Portpreserveparity   string `json:"portpreserveparity,omitempty"`
	Portpreserverange    string `json:"portpreserverange,omitempty"`
	Portquota            int    `json:"portquota,omitempty"`
	Sessionquota         int    `json:"sessionquota,omitempty"`
	Sessiontimeout       int    `json:"sessiontimeout,omitempty"`
	Stuntimeout          int    `json:"stuntimeout,omitempty"`
	Syncheck             string `json:"syncheck,omitempty"`
	Synidletimeout       int    `json:"synidletimeout,omitempty"`
	Transportprofilename string `json:"transportprofilename,omitempty"`
	Transportprotocol    string `json:"transportprotocol,omitempty"`
}
