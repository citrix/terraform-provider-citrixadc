package vpn

type Vpnglobalsslcertkeybinding struct {
	Cacert                 string `json:"cacert,omitempty"`
	Certkeyname            string `json:"certkeyname,omitempty"`
	Crlcheck               string `json:"crlcheck,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Ocspcheck              string `json:"ocspcheck,omitempty"`
	Userdataencryptionkey  string `json:"userdataencryptionkey,omitempty"`
}
