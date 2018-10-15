package lb

type Lbprofile struct {
	Cookiepassphrase              string `json:"cookiepassphrase,omitempty"`
	Dbslb                         string `json:"dbslb,omitempty"`
	Httponlycookieflag            string `json:"httponlycookieflag,omitempty"`
	Lbprofilename                 string `json:"lbprofilename,omitempty"`
	Processlocal                  string `json:"processlocal,omitempty"`
	Useencryptedpersistencecookie string `json:"useencryptedpersistencecookie,omitempty"`
	Usesecuredpersistencecookie   string `json:"usesecuredpersistencecookie,omitempty"`
	Vsvrcount                     int    `json:"vsvrcount,omitempty"`
}
