package transform

type Transformpolicy struct {
	Comment     string `json:"comment,omitempty"`
	Hits        int    `json:"hits,omitempty"`
	Isdefault   bool   `json:"isdefault,omitempty"`
	Logaction   string `json:"logaction,omitempty"`
	Name        string `json:"name,omitempty"`
	Newname     string `json:"newname,omitempty"`
	Profilename string `json:"profilename,omitempty"`
	Rule        string `json:"rule,omitempty"`
}
