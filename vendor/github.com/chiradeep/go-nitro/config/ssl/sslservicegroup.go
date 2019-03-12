package ssl

type Sslservicegroup struct {
	Ca                   bool   `json:"ca,omitempty"`
	Cipherredirect       string `json:"cipherredirect,omitempty"`
	Cipherurl            string `json:"cipherurl,omitempty"`
	Cleartextport        int    `json:"cleartextport,omitempty"`
	Clientauth           string `json:"clientauth,omitempty"`
	Clientcert           string `json:"clientcert,omitempty"`
	Commonname           string `json:"commonname,omitempty"`
	Crlcheck             string `json:"crlcheck,omitempty"`
	Dh                   string `json:"dh,omitempty"`
	Dhcount              int    `json:"dhcount,omitempty"`
	Dhfile               string `json:"dhfile,omitempty"`
	Dhkeyexpsizelimit    string `json:"dhkeyexpsizelimit,omitempty"`
	Ersa                 string `json:"ersa,omitempty"`
	Ersacount            int    `json:"ersacount,omitempty"`
	Nonfipsciphers       string `json:"nonfipsciphers,omitempty"`
	Ocspcheck            string `json:"ocspcheck,omitempty"`
	Ocspstapling         string `json:"ocspstapling,omitempty"`
	Redirectportrewrite  string `json:"redirectportrewrite,omitempty"`
	Sendclosenotify      string `json:"sendclosenotify,omitempty"`
	Serverauth           string `json:"serverauth,omitempty"`
	Servicegroupname     string `json:"servicegroupname,omitempty"`
	Servicename          string `json:"servicename,omitempty"`
	Sessreuse            string `json:"sessreuse,omitempty"`
	Sesstimeout          int    `json:"sesstimeout,omitempty"`
	Snicert              bool   `json:"snicert,omitempty"`
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
