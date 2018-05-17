package authentication

type Authenticationauthnprofile struct {
	Authenticationdomain string `json:"authenticationdomain,omitempty"`
	Authenticationhost   string `json:"authenticationhost,omitempty"`
	Authenticationlevel  int    `json:"authenticationlevel,omitempty"`
	Authnvsname          string `json:"authnvsname,omitempty"`
	Name                 string `json:"name,omitempty"`
}
