package authentication

type Authenticationcaptchaaction struct {
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Name                       string `json:"name,omitempty"`
	Secretkey                  string `json:"secretkey,omitempty"`
	Serverurl                  string `json:"serverurl,omitempty"`
	Sitekey                    string `json:"sitekey,omitempty"`
}
