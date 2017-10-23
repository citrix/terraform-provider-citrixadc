package stream

type Streamidentifier struct {
	Interval     int         `json:"interval,omitempty"`
	Name         string      `json:"name,omitempty"`
	Rule         interface{} `json:"rule,omitempty"`
	Samplecount  int         `json:"samplecount,omitempty"`
	Selectorname string      `json:"selectorname,omitempty"`
	Sort         string      `json:"sort,omitempty"`
}
