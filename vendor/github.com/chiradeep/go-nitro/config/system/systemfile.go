package system

type Systemfile struct {
	Fileaccesstime   string      `json:"fileaccesstime,omitempty"`
	Filecontent      string      `json:"filecontent,omitempty"`
	Fileencoding     string      `json:"fileencoding,omitempty"`
	Filelocation     string      `json:"filelocation,omitempty"`
	Filemode         interface{} `json:"filemode,omitempty"`
	Filemodifiedtime string      `json:"filemodifiedtime,omitempty"`
	Filename         string      `json:"filename,omitempty"`
	Filesize         int         `json:"filesize,omitempty"`
}
