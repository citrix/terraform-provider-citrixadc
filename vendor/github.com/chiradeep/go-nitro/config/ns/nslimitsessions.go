package ns

type Nslimitsessions struct {
	Detail          bool        `json:"detail,omitempty"`
	Drop            int         `json:"drop,omitempty"`
	Flag            int         `json:"flag,omitempty"`
	Flags           int         `json:"flags,omitempty"`
	Hits            int         `json:"hits,omitempty"`
	Limitidentifier string      `json:"limitidentifier,omitempty"`
	Maxbandwidth    int         `json:"maxbandwidth,omitempty"`
	Name            string      `json:"name,omitempty"`
	Number          interface{} `json:"number,omitempty"`
	Referencecount  int         `json:"referencecount,omitempty"`
	Selectoripv61   string      `json:"selectoripv61,omitempty"`
	Selectoripv62   string      `json:"selectoripv62,omitempty"`
	Timeout         int         `json:"timeout,omitempty"`
	Unit            int         `json:"unit,omitempty"`
}
