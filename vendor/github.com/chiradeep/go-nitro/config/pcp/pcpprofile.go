package pcp

type Pcpprofile struct {
	Announcemulticount int    `json:"announcemulticount,omitempty"`
	Mapping            string `json:"mapping,omitempty"`
	Maxmaplife         int    `json:"maxmaplife,omitempty"`
	Minmaplife         int    `json:"minmaplife,omitempty"`
	Name               string `json:"name,omitempty"`
	Peer               string `json:"peer,omitempty"`
	Thirdparty         string `json:"thirdparty,omitempty"`
}
