package aaa

type Aaakcdaccount struct {
	Cacert        string `json:"cacert,omitempty"`
	Delegateduser string `json:"delegateduser,omitempty"`
	Kcdaccount    string `json:"kcdaccount,omitempty"`
	Kcdpassword   string `json:"kcdpassword,omitempty"`
	Kcdspn        string `json:"kcdspn,omitempty"`
	Keytab        string `json:"keytab,omitempty"`
	Principle     string `json:"principle,omitempty"`
	Realmstr      string `json:"realmstr,omitempty"`
	Usercert      string `json:"usercert,omitempty"`
}
