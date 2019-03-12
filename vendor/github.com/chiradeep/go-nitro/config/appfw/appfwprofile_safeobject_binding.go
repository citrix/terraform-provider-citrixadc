package appfw

type Appfwprofilesafeobjectbinding struct {
	Action         interface{} `json:"action,omitempty"`
	Asexpression   string      `json:"as_expression,omitempty"`
	Comment        string      `json:"comment,omitempty"`
	Maxmatchlength int         `json:"maxmatchlength,omitempty"`
	Name           string      `json:"name,omitempty"`
	Safeobject     string      `json:"safeobject,omitempty"`
	State          string      `json:"state,omitempty"`
}
