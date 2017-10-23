package authentication

type Authenticationtacacsaction struct {
	Accounting                 string `json:"accounting,omitempty"`
	Auditfailedcmds            string `json:"auditfailedcmds,omitempty"`
	Authorization              string `json:"authorization,omitempty"`
	Authtimeout                int    `json:"authtimeout,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Failure                    int    `json:"failure,omitempty"`
	Name                       string `json:"name,omitempty"`
	Serverip                   string `json:"serverip,omitempty"`
	Serverport                 int    `json:"serverport,omitempty"`
	Success                    int    `json:"success,omitempty"`
	Tacacssecret               string `json:"tacacssecret,omitempty"`
}
