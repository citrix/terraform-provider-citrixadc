package appqoe

type Appqoepolicylbvserverbinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Bindpriority           int    `json:"bindpriority,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}
