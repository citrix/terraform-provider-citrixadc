package ns

type Nshttpprofile struct {
	Adpttimeout      string `json:"adpttimeout,omitempty"`
	Clientiphdrexpr  string `json:"clientiphdrexpr,omitempty"`
	Cmponpush        string `json:"cmponpush,omitempty"`
	Conmultiplex     string `json:"conmultiplex,omitempty"`
	Dropextracrlf    string `json:"dropextracrlf,omitempty"`
	Dropextradata    string `json:"dropextradata,omitempty"`
	Dropinvalreqs    string `json:"dropinvalreqs,omitempty"`
	Incomphdrdelay   int    `json:"incomphdrdelay,omitempty"`
	Markconnreqinval string `json:"markconnreqinval,omitempty"`
	Markhttp09inval  string `json:"markhttp09inval,omitempty"`
	Maxreq           int    `json:"maxreq,omitempty"`
	Maxreusepool     int    `json:"maxreusepool,omitempty"`
	Name             string `json:"name,omitempty"`
	Persistentetag   string `json:"persistentetag,omitempty"`
	Refcnt           int    `json:"refcnt,omitempty"`
	Reqtimeout       int    `json:"reqtimeout,omitempty"`
	Reqtimeoutaction string `json:"reqtimeoutaction,omitempty"`
	Spdy             string `json:"spdy,omitempty"`
	Weblog           string `json:"weblog,omitempty"`
	Websocket        string `json:"websocket,omitempty"`
}
