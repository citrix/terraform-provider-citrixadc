package basic

type Locationparameter struct {
	Context      string `json:"context,omitempty"`
	Custom       int    `json:"custom,omitempty"`
	Entries      int    `json:"entries,omitempty"`
	Errors       int    `json:"errors,omitempty"`
	Flags        int    `json:"flags,omitempty"`
	Format       string `json:"format,omitempty"`
	Lines        int    `json:"lines,omitempty"`
	Locationfile string `json:"Locationfile,omitempty"`
	Q1label      string `json:"q1label,omitempty"`
	Q2label      string `json:"q2label,omitempty"`
	Q3label      string `json:"q3label,omitempty"`
	Q4label      string `json:"q4label,omitempty"`
	Q5label      string `json:"q5label,omitempty"`
	Q6label      string `json:"q6label,omitempty"`
	Static       int    `json:"Static,omitempty"`
	Status       int    `json:"status,omitempty"`
	Warnings     int    `json:"warnings,omitempty"`
}
