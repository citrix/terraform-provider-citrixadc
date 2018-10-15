package authentication

type Authenticationstorefrontauthaction struct {
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Failure                    int    `json:"failure,omitempty"`
	Name                       string `json:"name,omitempty"`
	Serverurl                  string `json:"serverurl,omitempty"`
	Success                    int    `json:"success,omitempty"`
}
