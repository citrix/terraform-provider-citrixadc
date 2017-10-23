package appfw

type Appfwlearningdata struct {
	Cookieconsistency  string `json:"cookieconsistency,omitempty"`
	Crosssitescripting string `json:"crosssitescripting,omitempty"`
	Csrfformoriginurl  string `json:"csrfformoriginurl,omitempty"`
	Csrftag            string `json:"csrftag,omitempty"`
	Data               string `json:"data,omitempty"`
	Fieldconsistency   string `json:"fieldconsistency,omitempty"`
	Fieldformat        string `json:"fieldformat,omitempty"`
	Formactionurlff    string `json:"formactionurl_ff,omitempty"`
	Formactionurlffc   string `json:"formactionurl_ffc,omitempty"`
	Formactionurlsql   string `json:"formactionurl_sql,omitempty"`
	Formactionurlxss   string `json:"formactionurl_xss,omitempty"`
	Profilename        string `json:"profilename,omitempty"`
	Securitycheck      string `json:"securitycheck,omitempty"`
	Sqlinjection       string `json:"sqlinjection,omitempty"`
	Starturl           string `json:"starturl,omitempty"`
	Target             string `json:"target,omitempty"`
	Totalxmlrequests   bool   `json:"totalxmlrequests,omitempty"`
	Xmlattachmentcheck string `json:"xmlattachmentcheck,omitempty"`
	Xmldoscheck        string `json:"xmldoscheck,omitempty"`
	Xmlwsicheck        string `json:"xmlwsicheck,omitempty"`
}
