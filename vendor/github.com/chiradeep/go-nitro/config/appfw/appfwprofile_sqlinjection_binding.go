package appfw

type Appfwprofilesqlinjectionbinding struct {
	Asscanlocationsql string `json:"as_scan_location_sql,omitempty"`
	Comment           string `json:"comment,omitempty"`
	Formactionurlsql  string `json:"formactionurl_sql,omitempty"`
	Isregexsql        string `json:"isregex_sql,omitempty"`
	Name              string `json:"name,omitempty"`
	Sqlinjection      string `json:"sqlinjection,omitempty"`
	State             string `json:"state,omitempty"`
}
