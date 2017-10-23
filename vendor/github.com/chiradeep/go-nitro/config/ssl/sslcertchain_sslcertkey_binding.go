package ssl

type Sslcertchainsslcertkeybinding struct {
	Addsubject      bool   `json:"addsubject,omitempty"`
	Certkeyname     string `json:"certkeyname,omitempty"`
	Isca            bool   `json:"isca,omitempty"`
	Islinked        bool   `json:"islinked,omitempty"`
	Linkcertkeyname string `json:"linkcertkeyname,omitempty"`
}
