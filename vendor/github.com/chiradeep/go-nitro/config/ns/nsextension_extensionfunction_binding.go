package ns

type Nsextensionextensionfunctionbinding struct {
	Activeextensionfunction         int         `json:"activeextensionfunction,omitempty"`
	Extensionfuncdescription        string      `json:"extensionfuncdescription,omitempty"`
	Extensionfunctionallparams      interface{} `json:"extensionfunctionallparams,omitempty"`
	Extensionfunctionallparamscount int         `json:"extensionfunctionallparamscount,omitempty"`
	Extensionfunctionargcount       int         `json:"extensionfunctionargcount,omitempty"`
	Extensionfunctionargtype        interface{} `json:"extensionfunctionargtype,omitempty"`
	Extensionfunctionclasses        interface{} `json:"extensionfunctionclasses,omitempty"`
	Extensionfunctionclassescount   int         `json:"extensionfunctionclassescount,omitempty"`
	Extensionfunctionclasstype      string      `json:"extensionfunctionclasstype,omitempty"`
	Extensionfunctionlinenumber     int         `json:"extensionfunctionlinenumber,omitempty"`
	Extensionfunctionname           string      `json:"extensionfunctionname,omitempty"`
	Extensionfunctionreturntype     string      `json:"extensionfunctionreturntype,omitempty"`
	Name                            string      `json:"name,omitempty"`
}
