package appfw

type Appfwconfidfield struct {
	Comment   string `json:"comment,omitempty"`
	Fieldname string `json:"fieldname,omitempty"`
	Isregex   string `json:"isregex,omitempty"`
	State     string `json:"state,omitempty"`
	Url       string `json:"url,omitempty"`
}
