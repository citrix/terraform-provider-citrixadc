package ns

type Nstimer struct {
	Comment  string `json:"comment,omitempty"`
	Interval int    `json:"interval,omitempty"`
	Name     string `json:"name,omitempty"`
	Newname  string `json:"newname,omitempty"`
	Unit     string `json:"unit,omitempty"`
}
