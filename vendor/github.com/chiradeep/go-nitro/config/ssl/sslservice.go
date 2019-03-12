package ssl

type Sslservice struct {
	Cipherredirect       string `json:"cipherredirect,omitempty"`
	Cipherurl            string `json:"cipherurl,omitempty"`
	Clientauth           string `json:"clientauth,omitempty"`
	Clientcert           string `json:"clientcert,omitempty"`
	Commonname           string `json:"commonname,omitempty"`
	Dh                   string `json:"dh,omitempty"`
	Dhcount              int    `json:"dhcount,omitempty"`
	Dhfile               string `json:"dhfile,omitempty"`
	Dhkeyexpsizelimit    string `json:"dhkeyexpsizelimit,omitempty"`
	Dtlsflag             bool   `json:"dtlsflag,omitempty"`
	Dtlsprofilename      string `json:"dtlsprofilename,omitempty"`
	Ersa                 string `json:"ersa,omitempty"`
	Ersacount            int    `json:"ersacount,omitempty"`
	Nonfipsciphers       string `json:"nonfipsciphers,omitempty"`
	Ocspstapling         string `json:"ocspstapling,omitempty"`
	Pushenctrigger       string `json:"pushenctrigger,omitempty"`
	Redirectportrewrite  string `json:"redirectportrewrite,omitempty"`
	Sendclosenotify      string `json:"sendclosenotify,omitempty"`
	Serverauth           string `json:"serverauth,omitempty"`
	Service              int    `json:"service,omitempty"`
	Servicename          string `json:"servicename,omitempty"`
	Sessreuse            string `json:"sessreuse,omitempty"`
	Sesstimeout          int    `json:"sesstimeout,omitempty"`
	Skipcaname           bool   `json:"skipcaname,omitempty"`
	Snienable            string `json:"snienable,omitempty"`
	Ssl2                 string `json:"ssl2,omitempty"`
	Ssl3                 string `json:"ssl3,omitempty"`
	Sslprofile           string `json:"sslprofile,omitempty"`
	Sslredirect          string `json:"sslredirect,omitempty"`
	Sslv2redirect        string `json:"sslv2redirect,omitempty"`
	Sslv2url             string `json:"sslv2url,omitempty"`
	Strictsigdigestcheck string `json:"strictsigdigestcheck,omitempty"`
	Tls1                 string `json:"tls1,omitempty"`
	Tls11                string `json:"tls11,omitempty"`
	Tls12                string `json:"tls12,omitempty"`
	Tls13                string `json:"tls13,omitempty"`
}
