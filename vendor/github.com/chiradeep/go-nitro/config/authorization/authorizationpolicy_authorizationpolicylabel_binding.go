package authorization

type Authorizationpolicyauthorizationpolicylabelbinding struct {
	Boundto  string `json:"boundto,omitempty"`
	Name     string `json:"name,omitempty"`
	Priority int    `json:"priority,omitempty"`
}
