package utility

type Techsupport struct {
	Casenumber    string `json:"casenumber,omitempty"`
	Description   string `json:"description,omitempty"`
	File          string `json:"file,omitempty"`
	Partitionname string `json:"partitionname,omitempty"`
	Password      string `json:"password,omitempty"`
	Proxy         string `json:"proxy,omitempty"`
	Response      string `json:"response,omitempty"`
	Scope         string `json:"scope,omitempty"`
	Servername    string `json:"servername,omitempty"`
	Upload        bool   `json:"upload,omitempty"`
	Username      string `json:"username,omitempty"`
}
