package audit

type Auditmessages struct {
	Loglevel   interface{} `json:"loglevel,omitempty"`
	Numofmesgs int         `json:"numofmesgs,omitempty"`
	Value      string      `json:"value,omitempty"`
}
