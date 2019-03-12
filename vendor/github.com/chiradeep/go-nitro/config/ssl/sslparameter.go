package ssl

type Sslparameter struct {
	Crlmemorysizemb         int         `json:"crlmemorysizemb,omitempty"`
	Cryptodevdisablelimit   int         `json:"cryptodevdisablelimit,omitempty"`
	Defaultprofile          string      `json:"defaultprofile,omitempty"`
	Denysslreneg            string      `json:"denysslreneg,omitempty"`
	Dropreqwithnohostheader string      `json:"dropreqwithnohostheader,omitempty"`
	Encrypttriggerpktcount  int         `json:"encrypttriggerpktcount,omitempty"`
	Hybridfipsmode          string      `json:"hybridfipsmode,omitempty"`
	Insertcertspace         string      `json:"insertcertspace,omitempty"`
	Insertionencoding       string      `json:"insertionencoding,omitempty"`
	Montls1112disable       string      `json:"montls1112disable,omitempty"`
	Ocspcachesize           int         `json:"ocspcachesize,omitempty"`
	Pushenctriggertimeout   int         `json:"pushenctriggertimeout,omitempty"`
	Pushflag                int         `json:"pushflag,omitempty"`
	Quantumsize             string      `json:"quantumsize,omitempty"`
	Sendclosenotify         string      `json:"sendclosenotify,omitempty"`
	Sigdigesttype           interface{} `json:"sigdigesttype,omitempty"`
	Softwarecryptothreshold int         `json:"softwarecryptothreshold,omitempty"`
	Sslierrorcache          string      `json:"sslierrorcache,omitempty"`
	Sslimaxerrorcachemem    int         `json:"sslimaxerrorcachemem,omitempty"`
	Ssltriggertimeout       int         `json:"ssltriggertimeout,omitempty"`
	Strictcachecks          string      `json:"strictcachecks,omitempty"`
	Svctls1112disable       string      `json:"svctls1112disable,omitempty"`
	Undefactioncontrol      string      `json:"undefactioncontrol,omitempty"`
	Undefactiondata         string      `json:"undefactiondata,omitempty"`
}
