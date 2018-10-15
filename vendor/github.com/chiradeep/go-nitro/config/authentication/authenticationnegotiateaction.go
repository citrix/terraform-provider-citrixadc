package authentication

type Authenticationnegotiateaction struct {
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Domainuser                 string `json:"domainuser,omitempty"`
	Domainuserpasswd           string `json:"domainuserpasswd,omitempty"`
	Kcdspn                     string `json:"kcdspn,omitempty"`
	Keytab                     string `json:"keytab,omitempty"`
	Name                       string `json:"name,omitempty"`
	Ntlmpath                   string `json:"ntlmpath,omitempty"`
	Ou                         string `json:"ou,omitempty"`
}
