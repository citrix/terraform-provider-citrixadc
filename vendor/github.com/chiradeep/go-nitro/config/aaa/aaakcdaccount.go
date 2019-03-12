package aaa

type Aaakcdaccount struct {
	Cacert          string `json:"cacert,omitempty"`
	Delegateduser   string `json:"delegateduser,omitempty"`
	Enterpriserealm string `json:"enterpriserealm,omitempty"`
	Kcdaccount      string `json:"kcdaccount,omitempty"`
	Kcdpassword     string `json:"kcdpassword,omitempty"`
	Kcdspn          string `json:"kcdspn,omitempty"`
	Keytab          string `json:"keytab,omitempty"`
	Principle       string `json:"principle,omitempty"`
	Realmstr        string `json:"realmstr,omitempty"`
	Servicespn      string `json:"servicespn,omitempty"`
	Usercert        string `json:"usercert,omitempty"`
	Userrealm       string `json:"userrealm,omitempty"`
}
