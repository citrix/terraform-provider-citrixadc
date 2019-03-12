package appfw

type Appfwlearningdata struct {
	Asscanlocationsql      string `json:"as_scan_location_sql,omitempty"`
	Asscanlocationxss      string `json:"as_scan_location_xss,omitempty"`
	Asvalueexprsql         string `json:"as_value_expr_sql,omitempty"`
	Asvalueexprxss         string `json:"as_value_expr_xss,omitempty"`
	Asvaluetypesql         string `json:"as_value_type_sql,omitempty"`
	Asvaluetypexss         string `json:"as_value_type_xss,omitempty"`
	Contenttype            string `json:"contenttype,omitempty"`
	Cookieconsistency      string `json:"cookieconsistency,omitempty"`
	Creditcardnumber       string `json:"creditcardnumber,omitempty"`
	Creditcardnumberurl    string `json:"creditcardnumberurl,omitempty"`
	Crosssitescripting     string `json:"crosssitescripting,omitempty"`
	Csrfformoriginurl      string `json:"csrfformoriginurl,omitempty"`
	Csrftag                string `json:"csrftag,omitempty"`
	Data                   string `json:"data,omitempty"`
	Fieldconsistency       string `json:"fieldconsistency,omitempty"`
	Fieldformat            string `json:"fieldformat,omitempty"`
	Fieldformatcharmappcre string `json:"fieldformatcharmappcre,omitempty"`
	Fieldformatmaxlength   int    `json:"fieldformatmaxlength,omitempty"`
	Fieldformatminlength   int    `json:"fieldformatminlength,omitempty"`
	Fieldtype              string `json:"fieldtype,omitempty"`
	Formactionurlff        string `json:"formactionurl_ff,omitempty"`
	Formactionurlffc       string `json:"formactionurl_ffc,omitempty"`
	Formactionurlsql       string `json:"formactionurl_sql,omitempty"`
	Formactionurlxss       string `json:"formactionurl_xss,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Name                   string `json:"name,omitempty"`
	Profilename            string `json:"profilename,omitempty"`
	Securitycheck          string `json:"securitycheck,omitempty"`
	Sqlinjection           string `json:"sqlinjection,omitempty"`
	Starturl               string `json:"starturl,omitempty"`
	Target                 string `json:"target,omitempty"`
	Totalxmlrequests       bool   `json:"totalxmlrequests,omitempty"`
	Url                    string `json:"url,omitempty"`
	Value                  string `json:"value,omitempty"`
	Valuetype              string `json:"value_type,omitempty"`
	Xmlattachmentcheck     string `json:"xmlattachmentcheck,omitempty"`
	Xmldoscheck            string `json:"xmldoscheck,omitempty"`
	Xmlwsicheck            string `json:"xmlwsicheck,omitempty"`
}
