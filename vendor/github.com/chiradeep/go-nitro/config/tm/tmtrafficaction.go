package tm

type Tmtrafficaction struct {
	Apptimeout       int    `json:"apptimeout,omitempty"`
	Formssoaction    string `json:"formssoaction,omitempty"`
	Initiatelogout   string `json:"initiatelogout,omitempty"`
	Kcdaccount       string `json:"kcdaccount,omitempty"`
	Name             string `json:"name,omitempty"`
	Persistentcookie string `json:"persistentcookie,omitempty"`
	Samlssoprofile   string `json:"samlssoprofile,omitempty"`
	Sso              string `json:"sso,omitempty"`
}
