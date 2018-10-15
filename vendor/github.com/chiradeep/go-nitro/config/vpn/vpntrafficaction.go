package vpn

type Vpntrafficaction struct {
	Apptimeout       int    `json:"apptimeout,omitempty"`
	Formssoaction    string `json:"formssoaction,omitempty"`
	Fta              string `json:"fta,omitempty"`
	Hdx              string `json:"hdx,omitempty"`
	Kcdaccount       string `json:"kcdaccount,omitempty"`
	Name             string `json:"name,omitempty"`
	Passwdexpression string `json:"passwdexpression,omitempty"`
	Proxy            string `json:"proxy,omitempty"`
	Qual             string `json:"qual,omitempty"`
	Samlssoprofile   string `json:"samlssoprofile,omitempty"`
	Sso              string `json:"sso,omitempty"`
	Userexpression   string `json:"userexpression,omitempty"`
	Wanscaler        string `json:"wanscaler,omitempty"`
}
