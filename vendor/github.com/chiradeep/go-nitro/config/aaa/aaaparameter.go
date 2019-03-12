package aaa

type Aaaparameter struct {
	Aaadloglevel               string `json:"aaadloglevel,omitempty"`
	Aaadnatip                  string `json:"aaadnatip,omitempty"`
	Aaasessionloglevel         string `json:"aaasessionloglevel,omitempty"`
	Defaultauthtype            string `json:"defaultauthtype,omitempty"`
	Dynaddr                    string `json:"dynaddr,omitempty"`
	Enableenhancedauthfeedback string `json:"enableenhancedauthfeedback,omitempty"`
	Enablesessionstickiness    string `json:"enablesessionstickiness,omitempty"`
	Enablestaticpagecaching    string `json:"enablestaticpagecaching,omitempty"`
	Failedlogintimeout         int    `json:"failedlogintimeout,omitempty"`
	Ftmode                     string `json:"ftmode,omitempty"`
	Maxaaausers                int    `json:"maxaaausers,omitempty"`
	Maxkbquestions             int    `json:"maxkbquestions,omitempty"`
	Maxloginattempts           int    `json:"maxloginattempts,omitempty"`
	Maxsamldeflatesize         int    `json:"maxsamldeflatesize,omitempty"`
	Persistentloginattempts    string `json:"persistentloginattempts,omitempty"`
	Pwdexpirynotificationdays  int    `json:"pwdexpirynotificationdays,omitempty"`
}
