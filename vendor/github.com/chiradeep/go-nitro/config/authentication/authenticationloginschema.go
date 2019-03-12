package authentication

type Authenticationloginschema struct {
	Authenticationschema    string      `json:"authenticationschema,omitempty"`
	Authenticationstrength  int         `json:"authenticationstrength,omitempty"`
	Builtin                 interface{} `json:"builtin,omitempty"`
	Name                    string      `json:"name,omitempty"`
	Passwdexpression        string      `json:"passwdexpression,omitempty"`
	Passwordcredentialindex int         `json:"passwordcredentialindex,omitempty"`
	Ssocredentials          string      `json:"ssocredentials,omitempty"`
	Usercredentialindex     int         `json:"usercredentialindex,omitempty"`
	Userexpression          string      `json:"userexpression,omitempty"`
}
