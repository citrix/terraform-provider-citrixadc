package ns

type Nsextension struct {
	Comment           string `json:"comment,omitempty"`
	Detail            string `json:"detail,omitempty"`
	Functionhaltcount int    `json:"functionhaltcount,omitempty"`
	Functionhits      int    `json:"functionhits,omitempty"`
	Functionundefhits int    `json:"functionundefhits,omitempty"`
	Name              string `json:"name,omitempty"`
	Overwrite         bool   `json:"overwrite,omitempty"`
	Src               string `json:"src,omitempty"`
	Trace             string `json:"trace,omitempty"`
	Tracefunctions    string `json:"tracefunctions,omitempty"`
	Tracevariables    string `json:"tracevariables,omitempty"`
	Type              string `json:"type,omitempty"`
}
