package appfw

type Appfwprofilejsondosurlbinding struct {
	Alertonly                   string `json:"alertonly,omitempty"`
	Comment                     string `json:"comment,omitempty"`
	Isautodeployed              string `json:"isautodeployed,omitempty"`
	Jsondosurl                  string `json:"jsondosurl,omitempty"`
	Jsonmaxarraylength          int    `json:"jsonmaxarraylength,omitempty"`
	Jsonmaxarraylengthcheck     string `json:"jsonmaxarraylengthcheck,omitempty"`
	Jsonmaxcontainerdepth       int    `json:"jsonmaxcontainerdepth,omitempty"`
	Jsonmaxcontainerdepthcheck  string `json:"jsonmaxcontainerdepthcheck,omitempty"`
	Jsonmaxdocumentlength       int    `json:"jsonmaxdocumentlength,omitempty"`
	Jsonmaxdocumentlengthcheck  string `json:"jsonmaxdocumentlengthcheck,omitempty"`
	Jsonmaxobjectkeycount       int    `json:"jsonmaxobjectkeycount,omitempty"`
	Jsonmaxobjectkeycountcheck  string `json:"jsonmaxobjectkeycountcheck,omitempty"`
	Jsonmaxobjectkeylength      int    `json:"jsonmaxobjectkeylength,omitempty"`
	Jsonmaxobjectkeylengthcheck string `json:"jsonmaxobjectkeylengthcheck,omitempty"`
	Jsonmaxstringlength         int    `json:"jsonmaxstringlength,omitempty"`
	Jsonmaxstringlengthcheck    string `json:"jsonmaxstringlengthcheck,omitempty"`
	Name                        string `json:"name,omitempty"`
	State                       string `json:"state,omitempty"`
}
