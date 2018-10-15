package lsn

type Lsnhttphdrlogprofile struct {
	Httphdrlogprofilename string `json:"httphdrlogprofilename,omitempty"`
	Loghost               string `json:"loghost,omitempty"`
	Logmethod             string `json:"logmethod,omitempty"`
	Logurl                string `json:"logurl,omitempty"`
	Logversion            string `json:"logversion,omitempty"`
}
