package audit

type Auditsyslogaction struct {
	Acl                  string      `json:"acl,omitempty"`
	Alg                  string      `json:"alg,omitempty"`
	Appflowexport        string      `json:"appflowexport,omitempty"`
	Builtin              interface{} `json:"builtin,omitempty"`
	Contentinspectionlog string      `json:"contentinspectionlog,omitempty"`
	Dateformat           string      `json:"dateformat,omitempty"`
	Dns                  string      `json:"dns,omitempty"`
	Domainresolvenow     bool        `json:"domainresolvenow,omitempty"`
	Domainresolveretry   int         `json:"domainresolveretry,omitempty"`
	Feature              string      `json:"feature,omitempty"`
	Ip                   string      `json:"ip,omitempty"`
	Lbvservername        string      `json:"lbvservername,omitempty"`
	Logfacility          string      `json:"logfacility,omitempty"`
	Loglevel             interface{} `json:"loglevel,omitempty"`
	Lsn                  string      `json:"lsn,omitempty"`
	Maxlogdatasizetohold int         `json:"maxlogdatasizetohold,omitempty"`
	Name                 string      `json:"name,omitempty"`
	Netprofile           string      `json:"netprofile,omitempty"`
	Serverdomainname     string      `json:"serverdomainname,omitempty"`
	Serverip             string      `json:"serverip,omitempty"`
	Serverport           int         `json:"serverport,omitempty"`
	Sslinterception      string      `json:"sslinterception,omitempty"`
	Subscriberlog        string      `json:"subscriberlog,omitempty"`
	Tcp                  string      `json:"tcp,omitempty"`
	Tcpprofilename       string      `json:"tcpprofilename,omitempty"`
	Timezone             string      `json:"timezone,omitempty"`
	Transport            string      `json:"transport,omitempty"`
	Urlfiltering         string      `json:"urlfiltering,omitempty"`
	Userdefinedauditlog  string      `json:"userdefinedauditlog,omitempty"`
}
