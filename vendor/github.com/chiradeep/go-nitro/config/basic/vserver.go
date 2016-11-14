package basic

type Vserver struct {
	Backupvserver        string `json:"backupvserver,omitempty"`
	Cacheable            string `json:"cacheable,omitempty"`
	Clttimeout           int    `json:"clttimeout,omitempty"`
	Name                 string `json:"name,omitempty"`
	Pushvserver          string `json:"pushvserver,omitempty"`
	Redirecturl          string `json:"redirecturl,omitempty"`
	Somethod             string `json:"somethod,omitempty"`
	Sopersistence        string `json:"sopersistence,omitempty"`
	Sopersistencetimeout int    `json:"sopersistencetimeout,omitempty"`
	Sothreshold          int    `json:"sothreshold,omitempty"`
}
