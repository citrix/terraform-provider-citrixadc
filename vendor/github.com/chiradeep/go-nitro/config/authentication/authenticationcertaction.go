package authentication

type Authenticationcertaction struct {
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Groupnamefield             string `json:"groupnamefield,omitempty"`
	Name                       string `json:"name,omitempty"`
	Twofactor                  string `json:"twofactor,omitempty"`
	Usernamefield              string `json:"usernamefield,omitempty"`
}
