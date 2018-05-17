package wi

type Wisitetranslationinternalipbinding struct {
	Accesstype              string `json:"accesstype,omitempty"`
	Sitepath                string `json:"sitepath,omitempty"`
	Translationexternalip   string `json:"translationexternalip,omitempty"`
	Translationexternalport int    `json:"translationexternalport,omitempty"`
	Translationinternalip   string `json:"translationinternalip,omitempty"`
	Translationinternalport int    `json:"translationinternalport,omitempty"`
}
