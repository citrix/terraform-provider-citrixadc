package appflow

type Appflowparam struct {
	Aaausername        string `json:"aaausername,omitempty"`
	Appnamerefresh     int    `json:"appnamerefresh,omitempty"`
	Clienttrafficonly  string `json:"clienttrafficonly,omitempty"`
	Connectionchaining string `json:"connectionchaining,omitempty"`
	Flowrecordinterval int    `json:"flowrecordinterval,omitempty"`
	Httpauthorization  string `json:"httpauthorization,omitempty"`
	Httpcontenttype    string `json:"httpcontenttype,omitempty"`
	Httpcookie         string `json:"httpcookie,omitempty"`
	Httphost           string `json:"httphost,omitempty"`
	Httplocation       string `json:"httplocation,omitempty"`
	Httpmethod         string `json:"httpmethod,omitempty"`
	Httpreferer        string `json:"httpreferer,omitempty"`
	Httpsetcookie      string `json:"httpsetcookie,omitempty"`
	Httpsetcookie2     string `json:"httpsetcookie2,omitempty"`
	Httpurl            string `json:"httpurl,omitempty"`
	Httpuseragent      string `json:"httpuseragent,omitempty"`
	Httpvia            string `json:"httpvia,omitempty"`
	Httpxforwardedfor  string `json:"httpxforwardedfor,omitempty"`
	Templaterefresh    int    `json:"templaterefresh,omitempty"`
	Udppmtu            int    `json:"udppmtu,omitempty"`
}
