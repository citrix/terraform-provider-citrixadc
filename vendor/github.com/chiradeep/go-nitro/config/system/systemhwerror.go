package system

type Systemhwerror struct {
	Diskcheck    bool   `json:"diskcheck,omitempty"`
	Hwerrorcount int    `json:"hwerrorcount,omitempty"`
	Response     string `json:"response,omitempty"`
}
