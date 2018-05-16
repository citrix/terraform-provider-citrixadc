package ns

type Nsevents struct {
	Data0     int    `json:"data0,omitempty"`
	Data1     int    `json:"data1,omitempty"`
	Data2     int    `json:"data2,omitempty"`
	Data3     int    `json:"data3,omitempty"`
	Devid     int    `json:"devid,omitempty"`
	Devname   string `json:"devname,omitempty"`
	Eventcode int    `json:"eventcode,omitempty"`
	Eventno   int    `json:"eventno,omitempty"`
	Text      string `json:"text,omitempty"`
	Time      int    `json:"time,omitempty"`
}
