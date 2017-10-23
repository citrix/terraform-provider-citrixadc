package gslb

type Gslbvserverdomainbinding struct {
	Backupip         string `json:"backupip,omitempty"`
	Backupipflag     bool   `json:"backupipflag,omitempty"`
	Cookiedomain     string `json:"cookie_domain,omitempty"`
	Cookiedomainflag bool   `json:"cookie_domainflag,omitempty"`
	Cookietimeout    int    `json:"cookietimeout,omitempty"`
	Domainname       string `json:"domainname,omitempty"`
	Name             string `json:"name,omitempty"`
	Sitedomainttl    int    `json:"sitedomainttl,omitempty"`
	Ttl              int    `json:"ttl,omitempty"`
}
