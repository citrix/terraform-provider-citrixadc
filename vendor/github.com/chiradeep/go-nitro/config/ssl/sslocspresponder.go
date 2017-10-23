package ssl

type Sslocspresponder struct {
	Batchingdelay      int    `json:"batchingdelay,omitempty"`
	Batchingdepth      int    `json:"batchingdepth,omitempty"`
	Cache              string `json:"cache,omitempty"`
	Cachetimeout       int    `json:"cachetimeout,omitempty"`
	Dns                string `json:"dns,omitempty"`
	Insertclientcert   string `json:"insertclientcert,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Name               string `json:"name,omitempty"`
	Producedattimeskew int    `json:"producedattimeskew,omitempty"`
	Respondercert      string `json:"respondercert,omitempty"`
	Resptimeout        int    `json:"resptimeout,omitempty"`
	Signingcert        string `json:"signingcert,omitempty"`
	Trustresponder     bool   `json:"trustresponder,omitempty"`
	Url                string `json:"url,omitempty"`
	Useaia             bool   `json:"useaia,omitempty"`
	Usenonce           string `json:"usenonce,omitempty"`
}
