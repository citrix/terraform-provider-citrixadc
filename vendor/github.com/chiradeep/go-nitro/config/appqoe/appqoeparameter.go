package appqoe

type Appqoeparameter struct {
	Avgwaitingclient    int `json:"avgwaitingclient,omitempty"`
	Dosattackthresh     int `json:"dosattackthresh,omitempty"`
	Maxaltrespbandwidth int `json:"maxaltrespbandwidth,omitempty"`
	Sessionlife         int `json:"sessionlife,omitempty"`
}
