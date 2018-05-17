package authentication

type Authenticationdfaaction struct {
	Clientid                   string `json:"clientid,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Failure                    int    `json:"failure,omitempty"`
	Name                       string `json:"name,omitempty"`
	Passphrase                 string `json:"passphrase,omitempty"`
	Serverurl                  string `json:"serverurl,omitempty"`
	Success                    int    `json:"success,omitempty"`
}
