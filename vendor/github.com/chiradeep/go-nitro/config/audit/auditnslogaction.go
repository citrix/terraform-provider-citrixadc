package audit

type Auditnslogaction struct {
	Acl                 string      `json:"acl,omitempty"`
	Appflowexport       string      `json:"appflowexport,omitempty"`
	Builtin             interface{} `json:"builtin,omitempty"`
	Dateformat          string      `json:"dateformat,omitempty"`
	Logfacility         string      `json:"logfacility,omitempty"`
	Loglevel            interface{} `json:"loglevel,omitempty"`
	Name                string      `json:"name,omitempty"`
	Serverip            string      `json:"serverip,omitempty"`
	Serverport          int         `json:"serverport,omitempty"`
	Tcp                 string      `json:"tcp,omitempty"`
	Timezone            string      `json:"timezone,omitempty"`
	Userdefinedauditlog string      `json:"userdefinedauditlog,omitempty"`
}
