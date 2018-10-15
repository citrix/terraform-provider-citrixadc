package ns

type Nscapacity struct {
	Actualbandwidth int    `json:"actualbandwidth,omitempty"`
	Bandwidth       int    `json:"bandwidth,omitempty"`
	Edition         string `json:"edition,omitempty"`
	Instancecount   int    `json:"instancecount,omitempty"`
	Maxbandwidth    int    `json:"maxbandwidth,omitempty"`
	Maxvcpucount    int    `json:"maxvcpucount,omitempty"`
	Minbandwidth    int    `json:"minbandwidth,omitempty"`
	Nodeid          int    `json:"nodeid,omitempty"`
	Platform        string `json:"platform,omitempty"`
	Unit            string `json:"unit,omitempty"`
	Vcpu            bool   `json:"vcpu,omitempty"`
	Vcpucount       int    `json:"vcpucount,omitempty"`
}
