package wi

type Wisite struct {
	Agauthenticationmethod  string      `json:"agauthenticationmethod,omitempty"`
	Agcallbackurl           string      `json:"agcallbackurl,omitempty"`
	Agurl                   string      `json:"agurl,omitempty"`
	Appwelcomemessage       string      `json:"appwelcomemessage,omitempty"`
	Authenticationpoint     string      `json:"authenticationpoint,omitempty"`
	Defaultaccessmethod     string      `json:"defaultaccessmethod,omitempty"`
	Defaultcustomtextlocale string      `json:"defaultcustomtextlocale,omitempty"`
	Domainselection         string      `json:"domainselection,omitempty"`
	Footertext              string      `json:"footertext,omitempty"`
	Hidedomainfield         string      `json:"hidedomainfield,omitempty"`
	Kioskmode               string      `json:"kioskmode,omitempty"`
	Logindomains            string      `json:"logindomains,omitempty"`
	Loginsysmessage         string      `json:"loginsysmessage,omitempty"`
	Logintitle              string      `json:"logintitle,omitempty"`
	Preloginbutton          string      `json:"preloginbutton,omitempty"`
	Preloginmessage         string      `json:"preloginmessage,omitempty"`
	Prelogintitle           string      `json:"prelogintitle,omitempty"`
	Publishedresourcetype   string      `json:"publishedresourcetype,omitempty"`
	Restrictdomains         string      `json:"restrictdomains,omitempty"`
	Secondstaurl            string      `json:"secondstaurl,omitempty"`
	Sessionreliability      string      `json:"sessionreliability,omitempty"`
	Showrefresh             string      `json:"showrefresh,omitempty"`
	Showsearch              string      `json:"showsearch,omitempty"`
	Sitepath                string      `json:"sitepath,omitempty"`
	Sitetype                string      `json:"sitetype,omitempty"`
	Staurl                  string      `json:"staurl,omitempty"`
	Userinterfacebranding   string      `json:"userinterfacebranding,omitempty"`
	Userinterfacelayouts    string      `json:"userinterfacelayouts,omitempty"`
	Usetwotickets           string      `json:"usetwotickets,omitempty"`
	Websessiontimeout       int         `json:"websessiontimeout,omitempty"`
	Welcomemessage          string      `json:"welcomemessage,omitempty"`
	Wiauthenticationmethods interface{} `json:"wiauthenticationmethods,omitempty"`
	Wiuserinterfacemodes    string      `json:"wiuserinterfacemodes,omitempty"`
}
