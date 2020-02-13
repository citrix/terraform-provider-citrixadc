package system

type Systemrestorepoint struct {
	Backupfilename string `json:"backupfilename,omitempty"`
	Createdby      string `json:"createdby,omitempty"`
	Creationtime   string `json:"creationtime,omitempty"`
	Filename       string `json:"filename,omitempty"`
	Ipaddress      string `json:"ipaddress,omitempty"`
	Techsuprtname  string `json:"techsuprtname,omitempty"`
	Version        string `json:"version,omitempty"`
}
