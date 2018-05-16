package sc

type Scpolicy struct {
	Action            string `json:"action,omitempty"`
	Altcontentpath    string `json:"altcontentpath,omitempty"`
	Altcontentsvcname string `json:"altcontentsvcname,omitempty"`
	Delay             int    `json:"delay,omitempty"`
	Maxconn           int    `json:"maxconn,omitempty"`
	Name              string `json:"name,omitempty"`
	Rule              string `json:"rule,omitempty"`
	Url               string `json:"url,omitempty"`
}
