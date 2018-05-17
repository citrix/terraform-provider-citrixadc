package vpn

type Vpntrafficaction struct {
	Apptimeout     int    `json:"apptimeout,omitempty"`
	Formssoaction  string `json:"formssoaction,omitempty"`
	Fta            string `json:"fta,omitempty"`
	Kcdaccount     string `json:"kcdaccount,omitempty"`
	Name           string `json:"name,omitempty"`
	Qual           string `json:"qual,omitempty"`
	Samlssoprofile string `json:"samlssoprofile,omitempty"`
	Sso            string `json:"sso,omitempty"`
	Wanscaler      string `json:"wanscaler,omitempty"`
}
