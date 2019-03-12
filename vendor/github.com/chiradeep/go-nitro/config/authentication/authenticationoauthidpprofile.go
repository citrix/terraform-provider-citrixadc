package authentication

type Authenticationoauthidpprofile struct {
	Audience                   string `json:"audience,omitempty"`
	Clientid                   string `json:"clientid,omitempty"`
	Clientsecret               string `json:"clientsecret,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Encrypttoken               string `json:"encrypttoken,omitempty"`
	Issuer                     string `json:"issuer,omitempty"`
	Name                       string `json:"name,omitempty"`
	Oauthstatus                string `json:"oauthstatus,omitempty"`
	Redirecturl                string `json:"redirecturl,omitempty"`
	Refreshinterval            int    `json:"refreshinterval,omitempty"`
	Relyingpartymetadataurl    string `json:"relyingpartymetadataurl,omitempty"`
	Skewtime                   int    `json:"skewtime,omitempty"`
}
