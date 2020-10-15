package appfw

type Appfwprofilesafeobjectbinding struct {
	Action         interface{} `json:"action,omitempty"`
	Alertonly      string      `json:"alertonly,omitempty"`
	Asexpression   string      `json:"as_expression,omitempty"`
	Comment        string      `json:"comment,omitempty"`
	Isautodeployed string      `json:"isautodeployed,omitempty"`
	Maxmatchlength int         `json:"maxmatchlength,omitempty"`
	Name           string      `json:"name,omitempty"`
	Safeobject     string      `json:"safeobject,omitempty"`
	State          string      `json:"state,omitempty"`
}
