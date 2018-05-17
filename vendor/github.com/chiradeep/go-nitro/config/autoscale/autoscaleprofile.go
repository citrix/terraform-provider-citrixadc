package autoscale

type Autoscaleprofile struct {
	Apikey       string `json:"apikey,omitempty"`
	Name         string `json:"name,omitempty"`
	Sharedsecret string `json:"sharedsecret,omitempty"`
	Type         string `json:"type,omitempty"`
	Url          string `json:"url,omitempty"`
}
