package tm

type Tmsamlssoprofile struct {
	Assertionconsumerserviceurl string `json:"assertionconsumerserviceurl,omitempty"`
	Name                        string `json:"name,omitempty"`
	Relaystaterule              string `json:"relaystaterule,omitempty"`
	Samlissuername              string `json:"samlissuername,omitempty"`
	Samlsigningcertname         string `json:"samlsigningcertname,omitempty"`
	Sendpassword                string `json:"sendpassword,omitempty"`
}
