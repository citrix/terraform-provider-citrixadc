package network

type Mapbmr struct {
	Eabitlength    int    `json:"eabitlength,omitempty"`
	Name           string `json:"name,omitempty"`
	Psidlength     int    `json:"psidlength,omitempty"`
	Psidoffset     int    `json:"psidoffset,omitempty"`
	Ruleipv6prefix string `json:"ruleipv6prefix,omitempty"`
}
