package appfw

type Appfwprofilefieldconsistencybinding struct {
	Alertonly        string `json:"alertonly,omitempty"`
	Comment          string `json:"comment,omitempty"`
	Fieldconsistency string `json:"fieldconsistency,omitempty"`
	Formactionurlffc string `json:"formactionurl_ffc,omitempty"`
	Isautodeployed   string `json:"isautodeployed,omitempty"`
	Isregexffc       string `json:"isregex_ffc,omitempty"`
	Name             string `json:"name,omitempty"`
	State            string `json:"state,omitempty"`
}
