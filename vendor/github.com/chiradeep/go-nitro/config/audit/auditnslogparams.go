package audit

type Auditnslogparams struct {
	Acl                 string      `json:"acl,omitempty"`
	Alg                 string      `json:"alg,omitempty"`
	Appflowexport       string      `json:"appflowexport,omitempty"`
	Builtin             interface{} `json:"builtin,omitempty"`
	Dateformat          string      `json:"dateformat,omitempty"`
	Logfacility         string      `json:"logfacility,omitempty"`
	Loglevel            interface{} `json:"loglevel,omitempty"`
	Lsn                 string      `json:"lsn,omitempty"`
	Serverip            string      `json:"serverip,omitempty"`
	Serverport          int         `json:"serverport,omitempty"`
	Sslinterception     string      `json:"sslinterception,omitempty"`
	Subscriberlog       string      `json:"subscriberlog,omitempty"`
	Tcp                 string      `json:"tcp,omitempty"`
	Timezone            string      `json:"timezone,omitempty"`
	Urlfiltering        string      `json:"urlfiltering,omitempty"`
	Userdefinedauditlog string      `json:"userdefinedauditlog,omitempty"`
}
