package ns

type Nstimerautoscalepolicybinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Samplesize             int    `json:"samplesize,omitempty"`
	Threshold              int    `json:"threshold,omitempty"`
	Vserver                string `json:"vserver,omitempty"`
}
