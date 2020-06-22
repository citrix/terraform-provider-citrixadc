package policy

type Policyevaluation struct {
	Action                     string      `json:"action,omitempty"`
	Expression                 string      `json:"expression,omitempty"`
	Input                      string      `json:"input,omitempty"`
	Istruncatedrefresult       bool        `json:"istruncatedrefresult,omitempty"`
	Pitactionerrorresult       string      `json:"pitactionerrorresult,omitempty"`
	Pitactionevaltime          int         `json:"pitactionevaltime,omitempty"`
	Pitboolerrorresult         string      `json:"pitboolerrorresult,omitempty"`
	Pitboolevaltime            int         `json:"pitboolevaltime,omitempty"`
	Pitboolresult              bool        `json:"pitboolresult,omitempty"`
	Pitdoubleerrorresult       string      `json:"pitdoubleerrorresult,omitempty"`
	Pitdoubleevaltime          int         `json:"pitdoubleevaltime,omitempty"`
	Pitdoubleresult            float64     `json:"pitdoubleresult,omitempty"`
	Pitmodifiedinputdata       string      `json:"pitmodifiedinputdata,omitempty"`
	Pitnewoffsetarray          interface{} `json:"pitnewoffsetarray,omitempty"`
	Pitnumerrorresult          string      `json:"pitnumerrorresult,omitempty"`
	Pitnumevaltime             int         `json:"pitnumevaltime,omitempty"`
	Pitnumresult               int         `json:"pitnumresult,omitempty"`
	Pitoffseterrorresult       string      `json:"pitoffseterrorresult,omitempty"`
	Pitoffsetevaltime          int         `json:"pitoffsetevaltime,omitempty"`
	Pitoffsetlengtharray       interface{} `json:"pitoffsetlengtharray,omitempty"`
	Pitoffsetresult            int         `json:"pitoffsetresult,omitempty"`
	Pitoffsetresultlen         int         `json:"pitoffsetresultlen,omitempty"`
	Pitoldoffsetarray          interface{} `json:"pitoldoffsetarray,omitempty"`
	Pitoperationperformerarray interface{} `json:"pitoperationperformerarray,omitempty"`
	Pitreferrorresult          string      `json:"pitreferrorresult,omitempty"`
	Pitrefevaltime             int         `json:"pitrefevaltime,omitempty"`
	Pitrefresult               string      `json:"pitrefresult,omitempty"`
	Pitulongerrorresult        string      `json:"pitulongerrorresult,omitempty"`
	Pitulongevaltime           int         `json:"pitulongevaltime,omitempty"`
	Pitulongresult             int         `json:"pitulongresult,omitempty"`
	Type                       string      `json:"type,omitempty"`
}
