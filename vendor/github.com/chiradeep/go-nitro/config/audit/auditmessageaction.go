package audit

type Auditmessageaction struct {
	Bypasssafetycheck string `json:"bypasssafetycheck,omitempty"`
	Hits              int    `json:"hits,omitempty"`
	Loglevel          string `json:"loglevel,omitempty"`
	Loglevel1         string `json:"loglevel1,omitempty"`
	Logtonewnslog     string `json:"logtonewnslog,omitempty"`
	Name              string `json:"name,omitempty"`
	Referencecount    int    `json:"referencecount,omitempty"`
	Stringbuilderexpr string `json:"stringbuilderexpr,omitempty"`
	Undefhits         int    `json:"undefhits,omitempty"`
}
