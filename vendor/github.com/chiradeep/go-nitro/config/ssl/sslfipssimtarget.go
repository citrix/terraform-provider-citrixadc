package ssl

type Sslfipssimtarget struct {
	Certfile     string `json:"certfile,omitempty"`
	Keyvector    string `json:"keyvector,omitempty"`
	Sourcesecret string `json:"sourcesecret,omitempty"`
	Targetsecret string `json:"targetsecret,omitempty"`
}
