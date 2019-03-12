package vpn

type Vpnurl struct {
	Actualurl        string `json:"actualurl,omitempty"`
	Appjson          string `json:"appjson,omitempty"`
	Applicationtype  string `json:"applicationtype,omitempty"`
	Clientlessaccess string `json:"clientlessaccess,omitempty"`
	Comment          string `json:"comment,omitempty"`
	Iconurl          string `json:"iconurl,omitempty"`
	Linkname         string `json:"linkname,omitempty"`
	Samlssoprofile   string `json:"samlssoprofile,omitempty"`
	Ssotype          string `json:"ssotype,omitempty"`
	Urlname          string `json:"urlname,omitempty"`
	Vservername      string `json:"vservername,omitempty"`
}
