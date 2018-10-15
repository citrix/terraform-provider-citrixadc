package ica

type Icaaccessprofile struct {
	Builtin                    interface{} `json:"builtin,omitempty"`
	Clientaudioredirection     string      `json:"clientaudioredirection,omitempty"`
	Clientclipboardredirection string      `json:"clientclipboardredirection,omitempty"`
	Clientcomportredirection   string      `json:"clientcomportredirection,omitempty"`
	Clientdriveredirection     string      `json:"clientdriveredirection,omitempty"`
	Clientprinterredirection   string      `json:"clientprinterredirection,omitempty"`
	Clientusbdriveredirection  string      `json:"clientusbdriveredirection,omitempty"`
	Connectclientlptports      string      `json:"connectclientlptports,omitempty"`
	Isdefault                  bool        `json:"isdefault,omitempty"`
	Localremotedatasharing     string      `json:"localremotedatasharing,omitempty"`
	Multistream                string      `json:"multistream,omitempty"`
	Name                       string      `json:"name,omitempty"`
	Refcnt                     int         `json:"refcnt,omitempty"`
}
