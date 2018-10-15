package authentication

type Authenticationazurekeyvault struct {
	Clientid                   string `json:"clientid,omitempty"`
	Clientsecret               string `json:"clientsecret,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Name                       string `json:"name,omitempty"`
	Pushservice                string `json:"pushservice,omitempty"`
	Refreshinterval            int    `json:"refreshinterval,omitempty"`
	Servicekeyname             string `json:"servicekeyname,omitempty"`
	Signaturealg               string `json:"signaturealg,omitempty"`
	Tenantid                   string `json:"tenantid,omitempty"`
	Tokenendpoint              string `json:"tokenendpoint,omitempty"`
	Vaultname                  string `json:"vaultname,omitempty"`
}
