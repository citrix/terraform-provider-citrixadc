package ns

type Nsicapprofile struct {
	Allow204            string `json:"allow204,omitempty"`
	Connectionkeepalive string `json:"connectionkeepalive,omitempty"`
	Hostheader          string `json:"hostheader,omitempty"`
	Inserthttprequest   string `json:"inserthttprequest,omitempty"`
	Inserticapheaders   string `json:"inserticapheaders,omitempty"`
	Logaction           string `json:"logaction,omitempty"`
	Mode                string `json:"mode,omitempty"`
	Name                string `json:"name,omitempty"`
	Preview             string `json:"preview,omitempty"`
	Previewlength       int    `json:"previewlength,omitempty"`
	Queryparams         string `json:"queryparams,omitempty"`
	Reqtimeout          int    `json:"reqtimeout,omitempty"`
	Reqtimeoutaction    string `json:"reqtimeoutaction,omitempty"`
	Uri                 string `json:"uri,omitempty"`
	Useragent           string `json:"useragent,omitempty"`
}
