package wi

type Wisitefarmnamebinding struct {
	Farmname           string `json:"farmname,omitempty"`
	Groups             string `json:"groups,omitempty"`
	Loadbalance        string `json:"loadbalance,omitempty"`
	Recoveryfarm       string `json:"recoveryfarm,omitempty"`
	Sitepath           string `json:"sitepath,omitempty"`
	Sslrelayport       int    `json:"sslrelayport,omitempty"`
	Transport          string `json:"transport,omitempty"`
	Xmlport            int    `json:"xmlport,omitempty"`
	Xmlserveraddresses string `json:"xmlserveraddresses,omitempty"`
}
