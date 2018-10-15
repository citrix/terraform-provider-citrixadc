package ns

type Nshardware struct {
	Bmcrevision      string `json:"bmcrevision,omitempty"`
	Cpufrequncy      int    `json:"cpufrequncy,omitempty"`
	Encodedserialno  string `json:"encodedserialno,omitempty"`
	Host             string `json:"host,omitempty"`
	Hostid           int    `json:"hostid,omitempty"`
	Hwdescription    string `json:"hwdescription,omitempty"`
	Manufactureday   int    `json:"manufactureday,omitempty"`
	Manufacturemonth int    `json:"manufacturemonth,omitempty"`
	Manufactureyear  int    `json:"manufactureyear,omitempty"`
	Netscaleruuid    string `json:"netscaleruuid,omitempty"`
	Serialno         string `json:"serialno,omitempty"`
	Sysid            int    `json:"sysid,omitempty"`
}
