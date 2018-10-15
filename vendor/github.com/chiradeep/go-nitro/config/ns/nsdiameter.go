package ns

type Nsdiameter struct {
	Identity               string `json:"identity,omitempty"`
	Ownernode              int    `json:"ownernode,omitempty"`
	Realm                  string `json:"realm,omitempty"`
	Serverclosepropagation string `json:"serverclosepropagation,omitempty"`
}
