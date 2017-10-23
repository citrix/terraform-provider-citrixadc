package autoscale

type Autoscaleaction struct {
	Name                 string `json:"name,omitempty"`
	Parameters           string `json:"parameters,omitempty"`
	Profilename          string `json:"profilename,omitempty"`
	Quiettime            int    `json:"quiettime,omitempty"`
	Type                 string `json:"type,omitempty"`
	Vmdestroygraceperiod int    `json:"vmdestroygraceperiod,omitempty"`
	Vserver              string `json:"vserver,omitempty"`
}
