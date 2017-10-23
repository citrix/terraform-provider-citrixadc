package cs

type Csvserverdomainbinding struct {
	Appflowlog    string `json:"appflowlog,omitempty"`
	Backupip      string `json:"backupip,omitempty"`
	Cookiedomain  string `json:"cookiedomain,omitempty"`
	Cookietimeout int    `json:"cookietimeout,omitempty"`
	Domainname    string `json:"domainname,omitempty"`
	Name          string `json:"name,omitempty"`
	Sitedomainttl int    `json:"sitedomainttl,omitempty"`
	Ttl           int    `json:"ttl,omitempty"`
}
