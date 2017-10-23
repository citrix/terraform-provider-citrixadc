package wi

type Wisiteaccessmethodbinding struct {
	Accessmethod    string `json:"accessmethod,omitempty"`
	Clientipaddress string `json:"clientipaddress,omitempty"`
	Clientnetmask   string `json:"clientnetmask,omitempty"`
	Sitepath        string `json:"sitepath,omitempty"`
}
