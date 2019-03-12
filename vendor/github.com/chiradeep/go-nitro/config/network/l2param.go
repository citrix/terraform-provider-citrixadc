package network

type L2param struct {
	Bdggrpproxyarp         string `json:"bdggrpproxyarp,omitempty"`
	Bdgsetting             string `json:"bdgsetting,omitempty"`
	Bridgeagetimeout       int    `json:"bridgeagetimeout,omitempty"`
	Garponvridintf         string `json:"garponvridintf,omitempty"`
	Garpreply              string `json:"garpreply,omitempty"`
	Macmodefwdmypkt        string `json:"macmodefwdmypkt,omitempty"`
	Maxbridgecollision     int    `json:"maxbridgecollision,omitempty"`
	Mbfinstlearning        string `json:"mbfinstlearning,omitempty"`
	Mbfpeermacupdate       int    `json:"mbfpeermacupdate,omitempty"`
	Proxyarp               string `json:"proxyarp,omitempty"`
	Returntoethernetsender string `json:"returntoethernetsender,omitempty"`
	Rstintfonhafo          string `json:"rstintfonhafo,omitempty"`
	Skipproxyingbsdtraffic string `json:"skipproxyingbsdtraffic,omitempty"`
	Stopmacmoveupdate      string `json:"stopmacmoveupdate,omitempty"`
	Usemymac               string `json:"usemymac,omitempty"`
}
