package ns

type Nshardware struct {
	Cpufrequncy      int    `json:"cpufrequncy,omitempty"`
	Encodedserialno  string `json:"encodedserialno,omitempty"`
	Host             string `json:"host,omitempty"`
	Hostid           int    `json:"hostid,omitempty"`
	Hwdescription    string `json:"hwdescription,omitempty"`
	Manufactureday   int    `json:"manufactureday,omitempty"`
	Manufacturemonth int    `json:"manufacturemonth,omitempty"`
	Manufactureyear  int    `json:"manufactureyear,omitempty"`
	Serialno         string `json:"serialno,omitempty"`
	Sysid            int    `json:"sysid,omitempty"`
}
