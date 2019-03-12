package aaa

type Aaaldapparams struct {
	Authtimeout                int         `json:"authtimeout,omitempty"`
	Builtin                    interface{} `json:"builtin,omitempty"`
	Defaultauthenticationgroup string      `json:"defaultauthenticationgroup,omitempty"`
	Groupattrname              string      `json:"groupattrname,omitempty"`
	Groupauthname              string      `json:"groupauthname,omitempty"`
	Groupnameidentifier        string      `json:"groupnameidentifier,omitempty"`
	Groupsearchattribute       string      `json:"groupsearchattribute,omitempty"`
	Groupsearchfilter          string      `json:"groupsearchfilter,omitempty"`
	Groupsearchsubattribute    string      `json:"groupsearchsubattribute,omitempty"`
	Ldapbase                   string      `json:"ldapbase,omitempty"`
	Ldapbinddn                 string      `json:"ldapbinddn,omitempty"`
	Ldapbinddnpassword         string      `json:"ldapbinddnpassword,omitempty"`
	Ldaploginname              string      `json:"ldaploginname,omitempty"`
	Maxnestinglevel            int         `json:"maxnestinglevel,omitempty"`
	Nestedgroupextraction      string      `json:"nestedgroupextraction,omitempty"`
	Passwdchange               string      `json:"passwdchange,omitempty"`
	Searchfilter               string      `json:"searchfilter,omitempty"`
	Sectype                    string      `json:"sectype,omitempty"`
	Serverip                   string      `json:"serverip,omitempty"`
	Serverport                 int         `json:"serverport,omitempty"`
	Ssonameattribute           string      `json:"ssonameattribute,omitempty"`
	Subattributename           string      `json:"subattributename,omitempty"`
	Svrtype                    string      `json:"svrtype,omitempty"`
}
