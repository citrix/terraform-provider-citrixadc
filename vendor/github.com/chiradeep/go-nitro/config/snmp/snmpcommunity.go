package snmp

type Snmpcommunity struct {
	Communityname string `json:"communityname,omitempty"`
	Permissions   string `json:"permissions,omitempty"`
}
