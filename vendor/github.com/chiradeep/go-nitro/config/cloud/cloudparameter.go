package cloud

type Cloudparameter struct {
	Activationcode   string `json:"activationcode,omitempty"`
	Controllerfqdn   string `json:"controllerfqdn,omitempty"`
	Controllerport   int    `json:"controllerport,omitempty"`
	Customerid       string `json:"customerid,omitempty"`
	Instanceid       string `json:"instanceid,omitempty"`
	Resourcelocation string `json:"resourcelocation,omitempty"`
}
