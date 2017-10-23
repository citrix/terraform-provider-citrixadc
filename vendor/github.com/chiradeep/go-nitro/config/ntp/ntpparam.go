package ntp

type Ntpparam struct {
	Authentication string      `json:"authentication,omitempty"`
	Autokeylogsec  int         `json:"autokeylogsec,omitempty"`
	Revokelogsec   int         `json:"revokelogsec,omitempty"`
	Trustedkey     interface{} `json:"trustedkey,omitempty"`
}
