package cs

type Csvservercspolicybinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Cookieipport           string `json:"cookieipport,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Pipolicyhits           int    `json:"pipolicyhits,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Rule                   string `json:"rule,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
	Vserverid              string `json:"vserverid,omitempty"`
}
