package ssl

type Sslcert struct {
	Cacert         string `json:"cacert,omitempty"`
	Cacertform     string `json:"cacertform,omitempty"`
	Cakey          string `json:"cakey,omitempty"`
	Cakeyform      string `json:"cakeyform,omitempty"`
	Caserial       string `json:"caserial,omitempty"`
	Certfile       string `json:"certfile,omitempty"`
	Certform       string `json:"certform,omitempty"`
	Certtype       string `json:"certtype,omitempty"`
	Days           int    `json:"days,omitempty"`
	Keyfile        string `json:"keyfile,omitempty"`
	Keyform        string `json:"keyform,omitempty"`
	Pempassphrase  string `json:"pempassphrase,omitempty"`
	Reqfile        string `json:"reqfile,omitempty"`
	Subjectaltname string `json:"subjectaltname,omitempty"`
}
