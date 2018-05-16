package appflow

type Appflowaction struct {
	Collectors     interface{} `json:"collectors,omitempty"`
	Comment        string      `json:"comment,omitempty"`
	Description    string      `json:"description,omitempty"`
	Hits           int         `json:"hits,omitempty"`
	Name           string      `json:"name,omitempty"`
	Newname        string      `json:"newname,omitempty"`
	Referencecount int         `json:"referencecount,omitempty"`
}
