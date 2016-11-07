package basic

type Configstatus struct {
	Consistent            string `json:"consistent,omitempty"`
	Core                  int    `json:"core,omitempty"`
	Coreconfstring        string `json:"coreconfstring,omitempty"`
	Culpritcore           int    `json:"culpritcore,omitempty"`
	Culpritcoreconfstring string `json:"culpritcoreconfstring,omitempty"`
}
