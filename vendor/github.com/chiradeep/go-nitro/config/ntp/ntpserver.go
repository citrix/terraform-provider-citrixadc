package ntp

type Ntpserver struct {
	Autokey            bool   `json:"autokey,omitempty"`
	Key                int    `json:"key,omitempty"`
	Maxpoll            int    `json:"maxpoll,omitempty"`
	Minpoll            int    `json:"minpoll,omitempty"`
	Preferredntpserver string `json:"preferredntpserver,omitempty"`
	Serverip           string `json:"serverip,omitempty"`
	Servername         string `json:"servername,omitempty"`
}
