package gslb

type Gslbvservergslbservicebinding struct {
	Cnameentry         string `json:"cnameentry,omitempty"`
	Cumulativeweight   int    `json:"cumulativeweight,omitempty"`
	Curstate           string `json:"curstate,omitempty"`
	Domainname         string `json:"domainname,omitempty"`
	Dynamicconfwt      int    `json:"dynamicconfwt,omitempty"`
	Gslbboundsvctype   string `json:"gslbboundsvctype,omitempty"`
	Gslbthreshold      int    `json:"gslbthreshold,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Iscname            string `json:"iscname,omitempty"`
	Name               string `json:"name,omitempty"`
	Port               int    `json:"port,omitempty"`
	Preferredlocation  string `json:"preferredlocation,omitempty"`
	Servicename        string `json:"servicename,omitempty"`
	Sitepersistcookie  string `json:"sitepersistcookie,omitempty"`
	Svcsitepersistence string `json:"svcsitepersistence,omitempty"`
	Svreffgslbstate    string `json:"svreffgslbstate,omitempty"`
	Thresholdvalue     int    `json:"thresholdvalue,omitempty"`
	Weight             int    `json:"weight,omitempty"`
}
