package co

type Coparameter struct {
	Cachemaxage        bool `json:"cachemaxage,omitempty"`
	Imgtype            bool `json:"imgtype,omitempty"`
	Inlinecssthressize bool `json:"inlinecssthressize,omitempty"`
	Inlineimgthressize bool `json:"inlineimgthressize,omitempty"`
	Inlinejsthressize  bool `json:"inlinejsthressize,omitempty"`
	Jpegqualitypercent bool `json:"jpegqualitypercent,omitempty"`
}
