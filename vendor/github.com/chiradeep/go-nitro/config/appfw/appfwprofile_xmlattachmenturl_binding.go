package appfw

type Appfwprofilexmlattachmenturlbinding struct {
	Comment                       string `json:"comment,omitempty"`
	Name                          string `json:"name,omitempty"`
	State                         string `json:"state,omitempty"`
	Xmlattachmentcontenttype      string `json:"xmlattachmentcontenttype,omitempty"`
	Xmlattachmentcontenttypecheck string `json:"xmlattachmentcontenttypecheck,omitempty"`
	Xmlattachmenturl              string `json:"xmlattachmenturl,omitempty"`
	Xmlmaxattachmentsize          int    `json:"xmlmaxattachmentsize,omitempty"`
	Xmlmaxattachmentsizecheck     string `json:"xmlmaxattachmentsizecheck,omitempty"`
}
