package vpn

type Vpnintranetapplication struct {
	Clientapplication   interface{} `json:"clientapplication,omitempty"`
	Destip              string      `json:"destip,omitempty"`
	Destport            string      `json:"destport,omitempty"`
	Hostname            string      `json:"hostname,omitempty"`
	Interception        string      `json:"interception,omitempty"`
	Intranetapplication string      `json:"intranetapplication,omitempty"`
	Ipaddress           string      `json:"ipaddress,omitempty"`
	Iprange             string      `json:"iprange,omitempty"`
	Netmask             string      `json:"netmask,omitempty"`
	Protocol            string      `json:"protocol,omitempty"`
	Spoofiip            string      `json:"spoofiip,omitempty"`
	Srcip               string      `json:"srcip,omitempty"`
	Srcport             int         `json:"srcport,omitempty"`
}
