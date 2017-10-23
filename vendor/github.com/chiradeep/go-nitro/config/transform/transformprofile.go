package transform

type Transformprofile struct {
	Additionalreqheaderslist       string `json:"additionalreqheaderslist,omitempty"`
	Additionalrespheaderslist      string `json:"additionalrespheaderslist,omitempty"`
	Comment                        string `json:"comment,omitempty"`
	Name                           string `json:"name,omitempty"`
	Onlytransformabsurlinbody      string `json:"onlytransformabsurlinbody,omitempty"`
	Regexforfindingurlincss        string `json:"regexforfindingurlincss,omitempty"`
	Regexforfindingurlinjavascript string `json:"regexforfindingurlinjavascript,omitempty"`
	Regexforfindingurlinxcomponent string `json:"regexforfindingurlinxcomponent,omitempty"`
	Regexforfindingurlinxml        string `json:"regexforfindingurlinxml,omitempty"`
	Type                           string `json:"type,omitempty"`
}
