package appfw

type Appfwprofilesqlinjectionbinding struct {
	Alertonly         string `json:"alertonly,omitempty"`
	Asscanlocationsql string `json:"as_scan_location_sql,omitempty"`
	Asvalueexprsql    string `json:"as_value_expr_sql,omitempty"`
	Asvaluetypesql    string `json:"as_value_type_sql,omitempty"`
	Comment           string `json:"comment,omitempty"`
	Formactionurlsql  string `json:"formactionurl_sql,omitempty"`
	Isautodeployed    string `json:"isautodeployed,omitempty"`
	Isregexsql        string `json:"isregex_sql,omitempty"`
	Isvalueregexsql   string `json:"isvalueregex_sql,omitempty"`
	Name              string `json:"name,omitempty"`
	Sqlinjection      string `json:"sqlinjection,omitempty"`
	State             string `json:"state,omitempty"`
}
