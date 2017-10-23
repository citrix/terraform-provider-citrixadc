package system

type Systemglobalauthenticationradiuspolicybinding struct {
	Builtin    interface{} `json:"builtin,omitempty"`
	Policyname string      `json:"policyname,omitempty"`
	Priority   int         `json:"priority,omitempty"`
}
