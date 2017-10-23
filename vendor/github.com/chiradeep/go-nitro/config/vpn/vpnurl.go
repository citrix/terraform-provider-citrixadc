package vpn

type Vpnurl struct {
	Actualurl        string `json:"actualurl,omitempty"`
	Clientlessaccess string `json:"clientlessaccess,omitempty"`
	Comment          string `json:"comment,omitempty"`
	Linkname         string `json:"linkname,omitempty"`
	Urlname          string `json:"urlname,omitempty"`
}
