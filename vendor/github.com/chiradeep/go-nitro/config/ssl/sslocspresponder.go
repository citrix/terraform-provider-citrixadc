package ssl

type Sslocspresponder struct {
	Batchingdelay         int    `json:"batchingdelay,omitempty"`
	Batchingdepth         int    `json:"batchingdepth,omitempty"`
	Cache                 string `json:"cache,omitempty"`
	Cachetimeout          int    `json:"cachetimeout,omitempty"`
	Httpmethod            string `json:"httpmethod,omitempty"`
	Insertclientcert      string `json:"insertclientcert,omitempty"`
	Name                  string `json:"name,omitempty"`
	Ocspaiarefcount       int    `json:"ocspaiarefcount,omitempty"`
	Ocspipaddrstr         string `json:"ocspipaddrstr,omitempty"`
	Ocspurlresolvetimeout int    `json:"ocspurlresolvetimeout,omitempty"`
	Port                  int    `json:"port,omitempty"`
	Producedattimeskew    int    `json:"producedattimeskew,omitempty"`
	Respondercert         string `json:"respondercert,omitempty"`
	Resptimeout           int    `json:"resptimeout,omitempty"`
	Signingcert           string `json:"signingcert,omitempty"`
	Trustresponder        bool   `json:"trustresponder,omitempty"`
	Url                   string `json:"url,omitempty"`
	Usenonce              string `json:"usenonce,omitempty"`
}
