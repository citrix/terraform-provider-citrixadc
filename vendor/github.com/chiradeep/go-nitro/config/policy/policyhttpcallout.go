package policy

type Policyhttpcallout struct {
	Bodyexpr         string      `json:"bodyexpr,omitempty"`
	Cacheforsecs     int         `json:"cacheforsecs,omitempty"`
	Comment          string      `json:"comment,omitempty"`
	Effectivestate   string      `json:"effectivestate,omitempty"`
	Fullreqexpr      string      `json:"fullreqexpr,omitempty"`
	Headers          interface{} `json:"headers,omitempty"`
	Hits             int         `json:"hits,omitempty"`
	Hostexpr         string      `json:"hostexpr,omitempty"`
	Httpmethod       string      `json:"httpmethod,omitempty"`
	Ipaddress        string      `json:"ipaddress,omitempty"`
	Name             string      `json:"name,omitempty"`
	Parameters       interface{} `json:"parameters,omitempty"`
	Port             int         `json:"port,omitempty"`
	Recursivecallout int         `json:"recursivecallout,omitempty"`
	Resultexpr       string      `json:"resultexpr,omitempty"`
	Returntype       string      `json:"returntype,omitempty"`
	Scheme           string      `json:"scheme,omitempty"`
	Svrstate         string      `json:"svrstate,omitempty"`
	Undefhits        int         `json:"undefhits,omitempty"`
	Undefreason      string      `json:"undefreason,omitempty"`
	Urlstemexpr      string      `json:"urlstemexpr,omitempty"`
	Vserver          string      `json:"vserver,omitempty"`
}
