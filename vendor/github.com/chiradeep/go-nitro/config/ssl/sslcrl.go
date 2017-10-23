package ssl

type Sslcrl struct {
	Basedn           string `json:"basedn,omitempty"`
	Binary           string `json:"binary,omitempty"`
	Binddn           string `json:"binddn,omitempty"`
	Cacert           string `json:"cacert,omitempty"`
	Cacertfile       string `json:"cacertfile,omitempty"`
	Cakeyfile        string `json:"cakeyfile,omitempty"`
	Crlname          string `json:"crlname,omitempty"`
	Crlpath          string `json:"crlpath,omitempty"`
	Day              int    `json:"day,omitempty"`
	Daystoexpiration int    `json:"daystoexpiration,omitempty"`
	Flags            int    `json:"flags,omitempty"`
	Gencrl           string `json:"gencrl,omitempty"`
	Indexfile        string `json:"indexfile,omitempty"`
	Inform           string `json:"inform,omitempty"`
	Interval         string `json:"interval,omitempty"`
	Issuer           string `json:"issuer,omitempty"`
	Lastupdate       string `json:"lastupdate,omitempty"`
	Lastupdatetime   int    `json:"lastupdatetime,omitempty"`
	Method           string `json:"method,omitempty"`
	Nextupdate       string `json:"nextupdate,omitempty"`
	Password         string `json:"password,omitempty"`
	Port             int    `json:"port,omitempty"`
	Refresh          string `json:"refresh,omitempty"`
	Revoke           string `json:"revoke,omitempty"`
	Scope            string `json:"scope,omitempty"`
	Server           string `json:"server,omitempty"`
	Signaturealgo    string `json:"signaturealgo,omitempty"`
	Time             string `json:"time,omitempty"`
	Url              string `json:"url,omitempty"`
	Version          int    `json:"version,omitempty"`
}
