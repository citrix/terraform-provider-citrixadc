package ssl

type Sslcertreq struct {
	Challengepassword    string `json:"challengepassword,omitempty"`
	Commonname           string `json:"commonname,omitempty"`
	Companyname          string `json:"companyname,omitempty"`
	Countryname          string `json:"countryname,omitempty"`
	Digestmethod         string `json:"digestmethod,omitempty"`
	Emailaddress         string `json:"emailaddress,omitempty"`
	Fipskeyname          string `json:"fipskeyname,omitempty"`
	Keyfile              string `json:"keyfile,omitempty"`
	Keyform              string `json:"keyform,omitempty"`
	Localityname         string `json:"localityname,omitempty"`
	Organizationname     string `json:"organizationname,omitempty"`
	Organizationunitname string `json:"organizationunitname,omitempty"`
	Pempassphrase        string `json:"pempassphrase,omitempty"`
	Reqfile              string `json:"reqfile,omitempty"`
	Statename            string `json:"statename,omitempty"`
	Subjectaltname       string `json:"subjectaltname,omitempty"`
}
