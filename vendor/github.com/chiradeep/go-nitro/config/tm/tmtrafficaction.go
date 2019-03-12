package tm

type Tmtrafficaction struct {
	Apptimeout       int    `json:"apptimeout,omitempty"`
	Forcedtimeout    string `json:"forcedtimeout,omitempty"`
	Forcedtimeoutval int    `json:"forcedtimeoutval,omitempty"`
	Formssoaction    string `json:"formssoaction,omitempty"`
	Initiatelogout   string `json:"initiatelogout,omitempty"`
	Kcdaccount       string `json:"kcdaccount,omitempty"`
	Name             string `json:"name,omitempty"`
	Passwdexpression string `json:"passwdexpression,omitempty"`
	Persistentcookie string `json:"persistentcookie,omitempty"`
	Samlssoprofile   string `json:"samlssoprofile,omitempty"`
	Sso              string `json:"sso,omitempty"`
	Userexpression   string `json:"userexpression,omitempty"`
}
