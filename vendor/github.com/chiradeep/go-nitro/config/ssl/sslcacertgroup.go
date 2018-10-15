package ssl

type Sslcacertgroup struct {
	Cacertgroupname       string `json:"cacertgroupname,omitempty"`
	Cacertgroupreferences int    `json:"cacertgroupreferences,omitempty"`
	Crlcheck              string `json:"crlcheck,omitempty"`
	Ocspcheck             string `json:"ocspcheck,omitempty"`
}
