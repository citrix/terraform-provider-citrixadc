package authentication

type Authenticationldapaction struct {
	Authentication             string `json:"authentication,omitempty"`
	Authtimeout                int    `json:"authtimeout,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Failure                    int    `json:"failure,omitempty"`
	Followreferrals            string `json:"followreferrals,omitempty"`
	Groupattrname              string `json:"groupattrname,omitempty"`
	Groupnameidentifier        string `json:"groupnameidentifier,omitempty"`
	Groupsearchattribute       string `json:"groupsearchattribute,omitempty"`
	Groupsearchfilter          string `json:"groupsearchfilter,omitempty"`
	Groupsearchsubattribute    string `json:"groupsearchsubattribute,omitempty"`
	Ldapbase                   string `json:"ldapbase,omitempty"`
	Ldapbinddn                 string `json:"ldapbinddn,omitempty"`
	Ldapbinddnpassword         string `json:"ldapbinddnpassword,omitempty"`
	Ldaphostname               string `json:"ldaphostname,omitempty"`
	Ldaploginname              string `json:"ldaploginname,omitempty"`
	Maxldapreferrals           int    `json:"maxldapreferrals,omitempty"`
	Maxnestinglevel            int    `json:"maxnestinglevel,omitempty"`
	Name                       string `json:"name,omitempty"`
	Nestedgroupextraction      string `json:"nestedgroupextraction,omitempty"`
	Passwdchange               string `json:"passwdchange,omitempty"`
	Requireuser                string `json:"requireuser,omitempty"`
	Searchfilter               string `json:"searchfilter,omitempty"`
	Sectype                    string `json:"sectype,omitempty"`
	Serverip                   string `json:"serverip,omitempty"`
	Serverport                 int    `json:"serverport,omitempty"`
	Ssonameattribute           string `json:"ssonameattribute,omitempty"`
	Subattributename           string `json:"subattributename,omitempty"`
	Success                    int    `json:"success,omitempty"`
	Svrtype                    string `json:"svrtype,omitempty"`
	Validateservercert         string `json:"validateservercert,omitempty"`
}
