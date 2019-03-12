package contentinspection

type Contentinspectionaction struct {
	Hits             int    `json:"hits,omitempty"`
	Icapprofilename  string `json:"icapprofilename,omitempty"`
	Ifserverdown     string `json:"ifserverdown,omitempty"`
	Name             string `json:"name,omitempty"`
	Referencecount   int    `json:"referencecount,omitempty"`
	Reqtimeout       int    `json:"reqtimeout,omitempty"`
	Reqtimeoutaction string `json:"reqtimeoutaction,omitempty"`
	Serverip         string `json:"serverip,omitempty"`
	Servername       string `json:"servername,omitempty"`
	Serverport       int    `json:"serverport,omitempty"`
	Type             string `json:"type,omitempty"`
	Undefhits        int    `json:"undefhits,omitempty"`
}
