package cloud

type Cloudautoscalegroup struct {
	Azcount  int         `json:"azcount,omitempty"`
	Aznames  interface{} `json:"aznames,omitempty"`
	Graceful string      `json:"graceful,omitempty"`
	Name     string      `json:"name,omitempty"`
}
