package appfw

type Appfwprofilefileuploadtypebinding struct {
	Alertonly                 string      `json:"alertonly,omitempty"`
	Asfileuploadtypesurl      string      `json:"as_fileuploadtypes_url,omitempty"`
	Comment                   string      `json:"comment,omitempty"`
	Filetype                  interface{} `json:"filetype,omitempty"`
	Fileuploadtype            string      `json:"fileuploadtype,omitempty"`
	Isautodeployed            string      `json:"isautodeployed,omitempty"`
	Isregexfileuploadtypesurl string      `json:"isregex_fileuploadtypes_url,omitempty"`
	Name                      string      `json:"name,omitempty"`
	State                     string      `json:"state,omitempty"`
}
