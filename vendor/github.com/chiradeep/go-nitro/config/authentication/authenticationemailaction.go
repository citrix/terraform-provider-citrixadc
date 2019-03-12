package authentication

type Authenticationemailaction struct {
	Content                    string `json:"content,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Emailaddress               string `json:"emailaddress,omitempty"`
	Name                       string `json:"name,omitempty"`
	Password                   string `json:"password,omitempty"`
	Serverurl                  string `json:"serverurl,omitempty"`
	Timeout                    int    `json:"timeout,omitempty"`
	Type                       string `json:"type,omitempty"`
	Username                   string `json:"username,omitempty"`
}
