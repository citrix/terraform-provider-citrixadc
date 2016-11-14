package basic

type Server struct {
	Comment                   string `json:"comment,omitempty"`
	Delay                     int    `json:"delay,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Domainresolvenow          bool   `json:"domainresolvenow,omitempty"`
	Domainresolveretry        int    `json:"domainresolveretry,omitempty"`
	Graceful                  string `json:"graceful,omitempty"`
	Internal                  bool   `json:"Internal,omitempty"`
	Ipaddress                 string `json:"ipaddress,omitempty"`
	Ipv6address               string `json:"ipv6address,omitempty"`
	Name                      string `json:"name,omitempty"`
	Newname                   string `json:"newname,omitempty"`
	State                     string `json:"state,omitempty"`
	Statechangetimesec        string `json:"statechangetimesec,omitempty"`
	Td                        int    `json:"td,omitempty"`
	Tickssincelaststatechange int    `json:"tickssincelaststatechange,omitempty"`
	Translationip             string `json:"translationip,omitempty"`
	Translationmask           string `json:"translationmask,omitempty"`
}
