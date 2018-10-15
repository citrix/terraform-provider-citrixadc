package utility

type Install struct {
	Enhancedupgrade bool   `json:"enhancedupgrade,omitempty"`
	L               bool   `json:"l,omitempty"`
	Resizeswapvar   bool   `json:"resizeswapvar,omitempty"`
	Url             string `json:"url,omitempty"`
	Y               bool   `json:"y,omitempty"`
}
