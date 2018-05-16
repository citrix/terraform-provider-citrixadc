package appfw

type Appfwarchive struct {
	Comment  string `json:"comment,omitempty"`
	Name     string `json:"name,omitempty"`
	Response string `json:"response,omitempty"`
	Src      string `json:"src,omitempty"`
	Target   string `json:"target,omitempty"`
}
