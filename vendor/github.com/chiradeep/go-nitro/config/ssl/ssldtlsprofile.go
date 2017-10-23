package ssl

type Ssldtlsprofile struct {
	Helloverifyrequest string `json:"helloverifyrequest,omitempty"`
	Maxpacketsize      int    `json:"maxpacketsize,omitempty"`
	Maxrecordsize      int    `json:"maxrecordsize,omitempty"`
	Maxretrytime       int    `json:"maxretrytime,omitempty"`
	Name               string `json:"name,omitempty"`
	Pmtudiscovery      string `json:"pmtudiscovery,omitempty"`
	Terminatesession   string `json:"terminatesession,omitempty"`
}
