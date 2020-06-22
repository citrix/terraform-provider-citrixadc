package policy

type Policydataset struct {
	Comment     string `json:"comment,omitempty"`
	Description string `json:"description,omitempty"`
	Indextype   string `json:"indextype,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
}
