package vpn

type Vpnvservervpnepaprofilebinding struct {
	Acttype            int    `json:"acttype,omitempty"`
	Epaprofile         string `json:"epaprofile,omitempty"`
	Epaprofileoptional bool   `json:"epaprofileoptional,omitempty"`
	Name               string `json:"name,omitempty"`
}
