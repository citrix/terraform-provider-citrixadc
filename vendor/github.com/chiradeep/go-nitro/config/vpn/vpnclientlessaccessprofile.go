package vpn

type Vpnclientlessaccessprofile struct {
	Builtin                        interface{} `json:"builtin,omitempty"`
	Clientconsumedcookies          string      `json:"clientconsumedcookies,omitempty"`
	Cssrewritepolicylabel          string      `json:"cssrewritepolicylabel,omitempty"`
	Description                    string      `json:"description,omitempty"`
	Isdefault                      bool        `json:"isdefault,omitempty"`
	Javascriptrewritepolicylabel   string      `json:"javascriptrewritepolicylabel,omitempty"`
	Profilename                    string      `json:"profilename,omitempty"`
	Regexforfindingcustomurls      string      `json:"regexforfindingcustomurls,omitempty"`
	Regexforfindingurlincss        string      `json:"regexforfindingurlincss,omitempty"`
	Regexforfindingurlinjavascript string      `json:"regexforfindingurlinjavascript,omitempty"`
	Regexforfindingurlinxcomponent string      `json:"regexforfindingurlinxcomponent,omitempty"`
	Regexforfindingurlinxml        string      `json:"regexforfindingurlinxml,omitempty"`
	Reqhdrrewritepolicylabel       string      `json:"reqhdrrewritepolicylabel,omitempty"`
	Requirepersistentcookie        string      `json:"requirepersistentcookie,omitempty"`
	Reshdrrewritepolicylabel       string      `json:"reshdrrewritepolicylabel,omitempty"`
	Urlrewritepolicylabel          string      `json:"urlrewritepolicylabel,omitempty"`
	Xcomponentrewritepolicylabel   string      `json:"xcomponentrewritepolicylabel,omitempty"`
	Xmlrewritepolicylabel          string      `json:"xmlrewritepolicylabel,omitempty"`
}
