package rdp

type Rdpclientprofile struct {
	Addusernameinrdpfile string      `json:"addusernameinrdpfile,omitempty"`
	Audiocapturemode     string      `json:"audiocapturemode,omitempty"`
	Builtin              interface{} `json:"builtin,omitempty"`
	Keyboardhook         string      `json:"keyboardhook,omitempty"`
	Multimonitorsupport  string      `json:"multimonitorsupport,omitempty"`
	Name                 string      `json:"name,omitempty"`
	Psk                  string      `json:"psk,omitempty"`
	Randomizerdpfilename string      `json:"randomizerdpfilename,omitempty"`
	Rdpcookievalidity    int         `json:"rdpcookievalidity,omitempty"`
	Rdpcustomparams      string      `json:"rdpcustomparams,omitempty"`
	Rdpfilename          string      `json:"rdpfilename,omitempty"`
	Rdphost              string      `json:"rdphost,omitempty"`
	Rdplinkattribute     string      `json:"rdplinkattribute,omitempty"`
	Rdplistener          string      `json:"rdplistener,omitempty"`
	Rdpurloverride       string      `json:"rdpurloverride,omitempty"`
	Redirectclipboard    string      `json:"redirectclipboard,omitempty"`
	Redirectcomports     string      `json:"redirectcomports,omitempty"`
	Redirectdrives       string      `json:"redirectdrives,omitempty"`
	Redirectpnpdevices   string      `json:"redirectpnpdevices,omitempty"`
	Redirectprinters     string      `json:"redirectprinters,omitempty"`
	Videoplaybackmode    string      `json:"videoplaybackmode,omitempty"`
}
