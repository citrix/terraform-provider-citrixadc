package appfw

type Appfwprofilesafeobjectbinding struct {
	Action         interface{} `json:"action,omitempty"`
	Comment        string      `json:"comment,omitempty"`
	Expression     string      `json:"expression,omitempty"`
	Maxmatchlength int         `json:"maxmatchlength,omitempty"`
	Name           string      `json:"name,omitempty"`
	Safeobject     string      `json:"safeobject,omitempty"`
	State          string      `json:"state,omitempty"`
}
