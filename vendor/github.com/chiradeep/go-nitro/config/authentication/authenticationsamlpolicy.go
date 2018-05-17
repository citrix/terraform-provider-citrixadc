package authentication

type Authenticationsamlpolicy struct {
	Name      string `json:"name,omitempty"`
	Reqaction string `json:"reqaction,omitempty"`
	Rule      string `json:"rule,omitempty"`
}
