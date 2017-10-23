package cr

type Crvserverfilterpolicybinding struct {
	Name          string `json:"name,omitempty"`
	Policyname    string `json:"policyname,omitempty"`
	Priority      int    `json:"priority,omitempty"`
	Targetvserver string `json:"targetvserver,omitempty"`
}
