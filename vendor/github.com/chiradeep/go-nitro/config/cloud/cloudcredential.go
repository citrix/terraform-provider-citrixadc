package cloud

type Cloudcredential struct {
	Applicationid     string `json:"applicationid,omitempty"`
	Applicationsecret string `json:"applicationsecret,omitempty"`
	Isset             int    `json:"isset,omitempty"`
	Tenantidentifier  string `json:"tenantidentifier,omitempty"`
}
