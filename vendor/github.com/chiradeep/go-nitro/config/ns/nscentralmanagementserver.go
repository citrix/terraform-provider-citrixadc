package ns

type Nscentralmanagementserver struct {
	Ipaddress  string `json:"ipaddress,omitempty"`
	Password   string `json:"password,omitempty"`
	Servername string `json:"servername,omitempty"`
	Type       string `json:"type,omitempty"`
	Username   string `json:"username,omitempty"`
}
