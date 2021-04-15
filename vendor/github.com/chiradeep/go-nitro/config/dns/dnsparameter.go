package dns

type Dnsparameter struct {
	Builtin                 interface{} `json:"builtin,omitempty"`
	Cacheecszeroprefix      string      `json:"cacheecszeroprefix,omitempty"`
	Cachehitbypass          string      `json:"cachehitbypass,omitempty"`
	Cachenoexpire           string      `json:"cachenoexpire,omitempty"`
	Cacherecords            string      `json:"cacherecords,omitempty"`
	Dns64timeout            int         `json:"dns64timeout,omitempty"`
	Dnsrootreferral         string      `json:"dnsrootreferral,omitempty"`
	Dnssec                  string      `json:"dnssec,omitempty"`
	Ecsmaxsubnets           int         `json:"ecsmaxsubnets,omitempty"`
	Feature                 string      `json:"feature,omitempty"`
	Maxcachesize            int         `json:"maxcachesize,omitempty"`
	Maxnegativecachesize    int         `json:"maxnegativecachesize,omitempty"`
	Maxnegcachettl          int         `json:"maxnegcachettl,omitempty"`
	Maxpipeline             int         `json:"maxpipeline,omitempty"`
	Maxttl                  int         `json:"maxttl,omitempty"`
	Maxudppacketsize        int         `json:"maxudppacketsize,omitempty"`
	Minttl                  int         `json:"minttl,omitempty"`
	Namelookuppriority      string      `json:"namelookuppriority,omitempty"`
	Recursion               string      `json:"recursion,omitempty"`
	Resolutionorder         string      `json:"resolutionorder,omitempty"`
	Retries                 int         `json:"retries,omitempty"`
	Splitpktqueryprocessing string      `json:"splitpktqueryprocessing,omitempty"`
}
