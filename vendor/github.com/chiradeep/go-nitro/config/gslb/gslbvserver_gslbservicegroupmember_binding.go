package gslb

type Gslbvservergslbservicegroupmemberbinding struct {
	Curstate           string `json:"curstate,omitempty"`
	Dynamicweight      string `json:"dynamicweight,omitempty"`
	Gslbthreshold      int    `json:"gslbthreshold,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Name               string `json:"name,omitempty"`
	Port               int    `json:"port,omitempty"`
	Preferredlocation  string `json:"preferredlocation,omitempty"`
	Servicegroupname   string `json:"servicegroupname,omitempty"`
	Servicetype        string `json:"servicetype,omitempty"`
	Sitepersistcookie  string `json:"sitepersistcookie,omitempty"`
	Svcsitepersistence string `json:"svcsitepersistence,omitempty"`
	Svreffgslbstate    string `json:"svreffgslbstate,omitempty"`
	Thresholdvalue     int    `json:"thresholdvalue,omitempty"`
	Weight             int    `json:"weight,omitempty"`
}
