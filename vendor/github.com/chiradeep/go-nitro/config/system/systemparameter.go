package system

type Systemparameter struct {
	Allowdefaultpartition   string `json:"allowdefaultpartition,omitempty"`
	Basicauth               string `json:"basicauth,omitempty"`
	Cliloglevel             string `json:"cliloglevel,omitempty"`
	Doppler                 string `json:"doppler,omitempty"`
	Fipsusermode            string `json:"fipsusermode,omitempty"`
	Forcepasswordchange     string `json:"forcepasswordchange,omitempty"`
	Googleanalytics         string `json:"googleanalytics,omitempty"`
	Localauth               string `json:"localauth,omitempty"`
	Maxclient               int    `json:"maxclient,omitempty"`
	Minpasswordlen          int    `json:"minpasswordlen,omitempty"`
	Natpcbforceflushlimit   int    `json:"natpcbforceflushlimit,omitempty"`
	Natpcbrstontimeout      string `json:"natpcbrstontimeout,omitempty"`
	Promptstring            string `json:"promptstring,omitempty"`
	Rbaonresponse           string `json:"rbaonresponse,omitempty"`
	Reauthonauthparamchange string `json:"reauthonauthparamchange,omitempty"`
	Removesensitivefiles    string `json:"removesensitivefiles,omitempty"`
	Restrictedtimeout       string `json:"restrictedtimeout,omitempty"`
	Strongpassword          string `json:"strongpassword,omitempty"`
	Timeout                 int    `json:"timeout,omitempty"`
	Totalauthtimeout        int    `json:"totalauthtimeout,omitempty"`
}
