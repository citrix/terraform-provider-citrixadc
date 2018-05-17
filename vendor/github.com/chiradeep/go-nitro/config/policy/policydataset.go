package policy

type Policydataset struct {
	Description string `json:"description,omitempty"`
	Indextype   string `json:"indextype,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
}
