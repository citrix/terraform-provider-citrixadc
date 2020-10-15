package appfw

type Appfwprofilecsrftagbinding struct {
	Alertonly         string `json:"alertonly,omitempty"`
	Comment           string `json:"comment,omitempty"`
	Csrfformactionurl string `json:"csrfformactionurl,omitempty"`
	Csrftag           string `json:"csrftag,omitempty"`
	Isautodeployed    string `json:"isautodeployed,omitempty"`
	Name              string `json:"name,omitempty"`
	State             string `json:"state,omitempty"`
}
