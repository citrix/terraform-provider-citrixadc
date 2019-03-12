package authentication

type Authenticationepaaction struct {
	Csecexpr        string `json:"csecexpr,omitempty"`
	Defaultepagroup string `json:"defaultepagroup,omitempty"`
	Deletefiles     string `json:"deletefiles,omitempty"`
	Killprocess     string `json:"killprocess,omitempty"`
	Name            string `json:"name,omitempty"`
	Quarantinegroup string `json:"quarantinegroup,omitempty"`
}
