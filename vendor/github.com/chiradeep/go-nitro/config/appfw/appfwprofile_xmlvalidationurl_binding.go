package appfw

type Appfwprofilexmlvalidationurlbinding struct {
	Comment                  string `json:"comment,omitempty"`
	Name                     string `json:"name,omitempty"`
	State                    string `json:"state,omitempty"`
	Xmladditionalsoapheaders string `json:"xmladditionalsoapheaders,omitempty"`
	Xmlendpointcheck         string `json:"xmlendpointcheck,omitempty"`
	Xmlrequestschema         string `json:"xmlrequestschema,omitempty"`
	Xmlresponseschema        string `json:"xmlresponseschema,omitempty"`
	Xmlvalidateresponse      string `json:"xmlvalidateresponse,omitempty"`
	Xmlvalidatesoapenvelope  string `json:"xmlvalidatesoapenvelope,omitempty"`
	Xmlvalidationurl         string `json:"xmlvalidationurl,omitempty"`
	Xmlwsdl                  string `json:"xmlwsdl,omitempty"`
}
