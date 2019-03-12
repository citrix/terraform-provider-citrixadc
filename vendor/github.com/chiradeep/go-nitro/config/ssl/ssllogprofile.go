package ssl

type Ssllogprofile struct {
	Name                 string `json:"name,omitempty"`
	Ssllogclauth         string `json:"ssllogclauth,omitempty"`
	Ssllogclauthfailures string `json:"ssllogclauthfailures,omitempty"`
	Sslloghs             string `json:"sslloghs,omitempty"`
	Sslloghsfailures     string `json:"sslloghsfailures,omitempty"`
}
