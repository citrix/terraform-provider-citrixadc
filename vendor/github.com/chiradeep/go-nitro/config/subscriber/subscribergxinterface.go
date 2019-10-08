package subscriber

type Subscribergxinterface struct {
	Cerrequesttimeout         int         `json:"cerrequesttimeout,omitempty"`
	Gxreportingavp1           interface{} `json:"gxreportingavp1,omitempty"`
	Gxreportingavp1type       string      `json:"gxreportingavp1type,omitempty"`
	Gxreportingavp1vendorid   int         `json:"gxreportingavp1vendorid,omitempty"`
	Gxreportingavp2           interface{} `json:"gxreportingavp2,omitempty"`
	Gxreportingavp2type       string      `json:"gxreportingavp2type,omitempty"`
	Gxreportingavp2vendorid   int         `json:"gxreportingavp2vendorid,omitempty"`
	Gxreportingavp3           interface{} `json:"gxreportingavp3,omitempty"`
	Gxreportingavp3type       string      `json:"gxreportingavp3type,omitempty"`
	Gxreportingavp3vendorid   int         `json:"gxreportingavp3vendorid,omitempty"`
	Gxreportingavp4           interface{} `json:"gxreportingavp4,omitempty"`
	Gxreportingavp4type       string      `json:"gxreportingavp4type,omitempty"`
	Gxreportingavp4vendorid   int         `json:"gxreportingavp4vendorid,omitempty"`
	Gxreportingavp5           interface{} `json:"gxreportingavp5,omitempty"`
	Gxreportingavp5type       string      `json:"gxreportingavp5type,omitempty"`
	Gxreportingavp5vendorid   int         `json:"gxreportingavp5vendorid,omitempty"`
	Healthcheck               string      `json:"healthcheck,omitempty"`
	Healthcheckttl            int         `json:"healthcheckttl,omitempty"`
	Holdonsubscriberabsence   string      `json:"holdonsubscriberabsence,omitempty"`
	Identity                  string      `json:"identity,omitempty"`
	Idlettl                   int         `json:"idlettl,omitempty"`
	Negativettl               int         `json:"negativettl,omitempty"`
	Negativettllimitedsuccess string      `json:"negativettllimitedsuccess,omitempty"`
	Nodeid                    int         `json:"nodeid,omitempty"`
	Pcrfrealm                 string      `json:"pcrfrealm,omitempty"`
	Purgesdbongxfailure       string      `json:"purgesdbongxfailure,omitempty"`
	Realm                     string      `json:"realm,omitempty"`
	Requestretryattempts      int         `json:"requestretryattempts,omitempty"`
	Requesttimeout            int         `json:"requesttimeout,omitempty"`
	Revalidationtimeout       int         `json:"revalidationtimeout,omitempty"`
	Service                   string      `json:"service,omitempty"`
	Servicepathavp            interface{} `json:"servicepathavp,omitempty"`
	Servicepathinfomode       string      `json:"servicepathinfomode,omitempty"`
	Servicepathvendorid       int         `json:"servicepathvendorid,omitempty"`
	Status                    string      `json:"status,omitempty"`
	Svrstate                  string      `json:"svrstate,omitempty"`
	Vserver                   string      `json:"vserver,omitempty"`
}
