package ns

type Nshttpprofile struct {
	Adpttimeout               string      `json:"adpttimeout,omitempty"`
	Altsvc                    string      `json:"altsvc,omitempty"`
	Apdexcltresptimethreshold int         `json:"apdexcltresptimethreshold,omitempty"`
	Apdexsvrresptimethreshold int         `json:"apdexsvrresptimethreshold,omitempty"`
	Builtin                   interface{} `json:"builtin,omitempty"`
	Clientiphdrexpr           string      `json:"clientiphdrexpr,omitempty"`
	Cmponpush                 string      `json:"cmponpush,omitempty"`
	Conmultiplex              string      `json:"conmultiplex,omitempty"`
	Dropextracrlf             string      `json:"dropextracrlf,omitempty"`
	Dropextradata             string      `json:"dropextradata,omitempty"`
	Dropinvalreqs             string      `json:"dropinvalreqs,omitempty"`
	Feature                   string      `json:"feature,omitempty"`
	Http2                     string      `json:"http2,omitempty"`
	Http2direct               string      `json:"http2direct,omitempty"`
	Http2headertablesize      int         `json:"http2headertablesize,omitempty"`
	Http2initialwindowsize    int         `json:"http2initialwindowsize,omitempty"`
	Http2maxconcurrentstreams int         `json:"http2maxconcurrentstreams,omitempty"`
	Http2maxframesize         int         `json:"http2maxframesize,omitempty"`
	Http2maxheaderlistsize    int         `json:"http2maxheaderlistsize,omitempty"`
	Http2minseverconn         int         `json:"http2minseverconn,omitempty"`
	Http2strictcipher         string      `json:"http2strictcipher,omitempty"`
	Incomphdrdelay            int         `json:"incomphdrdelay,omitempty"`
	Markconnreqinval          string      `json:"markconnreqinval,omitempty"`
	Markhttp09inval           string      `json:"markhttp09inval,omitempty"`
	Marktracereqinval         string      `json:"marktracereqinval,omitempty"`
	Maxheaderlen              int         `json:"maxheaderlen,omitempty"`
	Maxreq                    int         `json:"maxreq,omitempty"`
	Maxreusepool              int         `json:"maxreusepool,omitempty"`
	Minreusepool              int         `json:"minreusepool,omitempty"`
	Name                      string      `json:"name,omitempty"`
	Persistentetag            string      `json:"persistentetag,omitempty"`
	Refcnt                    int         `json:"refcnt,omitempty"`
	Reqtimeout                int         `json:"reqtimeout,omitempty"`
	Reqtimeoutaction          string      `json:"reqtimeoutaction,omitempty"`
	Reusepooltimeout          int         `json:"reusepooltimeout,omitempty"`
	Rtsptunnel                string      `json:"rtsptunnel,omitempty"`
	Spdy                      string      `json:"spdy,omitempty"`
	Weblog                    string      `json:"weblog,omitempty"`
	Websocket                 string      `json:"websocket,omitempty"`
}
