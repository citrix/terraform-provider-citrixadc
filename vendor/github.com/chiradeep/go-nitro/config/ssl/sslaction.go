package ssl

type Sslaction struct {
	Builtin                interface{} `json:"builtin,omitempty"`
	Certhashheader         string      `json:"certhashheader,omitempty"`
	Certheader             string      `json:"certheader,omitempty"`
	Certissuerheader       string      `json:"certissuerheader,omitempty"`
	Certnotafterheader     string      `json:"certnotafterheader,omitempty"`
	Certnotbeforeheader    string      `json:"certnotbeforeheader,omitempty"`
	Certserialheader       string      `json:"certserialheader,omitempty"`
	Certsubjectheader      string      `json:"certsubjectheader,omitempty"`
	Cipher                 string      `json:"cipher,omitempty"`
	Cipherheader           string      `json:"cipherheader,omitempty"`
	Clientauth             string      `json:"clientauth,omitempty"`
	Clientcert             string      `json:"clientcert,omitempty"`
	Clientcerthash         string      `json:"clientcerthash,omitempty"`
	Clientcertissuer       string      `json:"clientcertissuer,omitempty"`
	Clientcertnotafter     string      `json:"clientcertnotafter,omitempty"`
	Clientcertnotbefore    string      `json:"clientcertnotbefore,omitempty"`
	Clientcertserialnumber string      `json:"clientcertserialnumber,omitempty"`
	Clientcertsubject      string      `json:"clientcertsubject,omitempty"`
	Description            string      `json:"description,omitempty"`
	Hits                   int         `json:"hits,omitempty"`
	Name                   string      `json:"name,omitempty"`
	Owasupport             string      `json:"owasupport,omitempty"`
	Referencecount         int         `json:"referencecount,omitempty"`
	Sessionid              string      `json:"sessionid,omitempty"`
	Sessionidheader        string      `json:"sessionidheader,omitempty"`
	Undefhits              int         `json:"undefhits,omitempty"`
}
