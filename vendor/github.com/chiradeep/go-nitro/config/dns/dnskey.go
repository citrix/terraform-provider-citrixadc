package dns

type Dnskey struct {
	Algorithm          string `json:"algorithm,omitempty"`
	Expires            int    `json:"expires,omitempty"`
	Filenameprefix     string `json:"filenameprefix,omitempty"`
	Keyname            string `json:"keyname,omitempty"`
	Keysize            int    `json:"keysize,omitempty"`
	Keytype            string `json:"keytype,omitempty"`
	Notificationperiod int    `json:"notificationperiod,omitempty"`
	Password           string `json:"password,omitempty"`
	Privatekey         string `json:"privatekey,omitempty"`
	Publickey          string `json:"publickey,omitempty"`
	Src                string `json:"src,omitempty"`
	Ttl                int    `json:"ttl,omitempty"`
	Units1             string `json:"units1,omitempty"`
	Units2             string `json:"units2,omitempty"`
	Zonename           string `json:"zonename,omitempty"`
}
