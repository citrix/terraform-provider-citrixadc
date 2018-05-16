package appfw

type Appfwxmlcontenttype struct {
	Builtin             interface{} `json:"builtin,omitempty"`
	Isregex             string      `json:"isregex,omitempty"`
	Xmlcontenttypevalue string      `json:"xmlcontenttypevalue,omitempty"`
}
