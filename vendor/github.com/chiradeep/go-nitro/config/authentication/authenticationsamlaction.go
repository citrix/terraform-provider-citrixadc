package authentication

type Authenticationsamlaction struct {
	Defaultauthenticationgroup  string `json:"defaultauthenticationgroup,omitempty"`
	Name                        string `json:"name,omitempty"`
	Samlidpcertname             string `json:"samlidpcertname,omitempty"`
	Samlissuername              string `json:"samlissuername,omitempty"`
	Samlredirecturl             string `json:"samlredirecturl,omitempty"`
	Samlrejectunsignedassertion string `json:"samlrejectunsignedassertion,omitempty"`
	Samlsigningcertname         string `json:"samlsigningcertname,omitempty"`
	Samltwofactor               string `json:"samltwofactor,omitempty"`
	Samluserfield               string `json:"samluserfield,omitempty"`
}
