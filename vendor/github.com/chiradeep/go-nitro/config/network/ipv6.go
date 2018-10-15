package network

type Ipv6 struct {
	Basereachtime        int    `json:"basereachtime,omitempty"`
	Dodad                string `json:"dodad,omitempty"`
	Natprefix            string `json:"natprefix,omitempty"`
	Ndbasereachtime      int    `json:"ndbasereachtime,omitempty"`
	Ndreachtime          int    `json:"ndreachtime,omitempty"`
	Ndretransmissiontime int    `json:"ndretransmissiontime,omitempty"`
	Ralearning           string `json:"ralearning,omitempty"`
	Reachtime            int    `json:"reachtime,omitempty"`
	Retransmissiontime   int    `json:"retransmissiontime,omitempty"`
	Routerredirection    string `json:"routerredirection,omitempty"`
	Td                   int    `json:"td,omitempty"`
	Usipnatprefix        string `json:"usipnatprefix,omitempty"`
}
