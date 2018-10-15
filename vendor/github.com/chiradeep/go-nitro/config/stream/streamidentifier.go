package stream

type Streamidentifier struct {
	Acceptancethreshold     string      `json:"acceptancethreshold,omitempty"`
	Appflowlog              string      `json:"appflowlog,omitempty"`
	Breachthreshold         int         `json:"breachthreshold,omitempty"`
	Interval                int         `json:"interval,omitempty"`
	Maxtransactionthreshold int         `json:"maxtransactionthreshold,omitempty"`
	Mintransactionthreshold int         `json:"mintransactionthreshold,omitempty"`
	Name                    string      `json:"name,omitempty"`
	Rule                    interface{} `json:"rule,omitempty"`
	Samplecount             int         `json:"samplecount,omitempty"`
	Selectorname            string      `json:"selectorname,omitempty"`
	Snmptrap                string      `json:"snmptrap,omitempty"`
	Sort                    string      `json:"sort,omitempty"`
	Trackackonlypackets     string      `json:"trackackonlypackets,omitempty"`
	Tracktransactions       string      `json:"tracktransactions,omitempty"`
}
