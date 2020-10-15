package appfw

type Appfwjsoncontenttype struct {
	Builtin              interface{} `json:"builtin,omitempty"`
	Feature              string      `json:"feature,omitempty"`
	Isregex              string      `json:"isregex,omitempty"`
	Jsoncontenttypevalue string      `json:"jsoncontenttypevalue,omitempty"`
}
