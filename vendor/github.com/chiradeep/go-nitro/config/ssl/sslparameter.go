package ssl

type Sslparameter struct {
	Crlmemorysizemb         int    `json:"crlmemorysizemb,omitempty"`
	Denysslreneg            string `json:"denysslreneg,omitempty"`
	Dropreqwithnohostheader string `json:"dropreqwithnohostheader,omitempty"`
	Encrypttriggerpktcount  int    `json:"encrypttriggerpktcount,omitempty"`
	Insertionencoding       string `json:"insertionencoding,omitempty"`
	Ocspcachesize           int    `json:"ocspcachesize,omitempty"`
	Pushenctriggertimeout   int    `json:"pushenctriggertimeout,omitempty"`
	Pushflag                int    `json:"pushflag,omitempty"`
	Quantumsize             string `json:"quantumsize,omitempty"`
	Sendclosenotify         string `json:"sendclosenotify,omitempty"`
	Ssltriggertimeout       int    `json:"ssltriggertimeout,omitempty"`
	Strictcachecks          string `json:"strictcachecks,omitempty"`
	Undefactioncontrol      string `json:"undefactioncontrol,omitempty"`
	Undefactiondata         string `json:"undefactiondata,omitempty"`
}
