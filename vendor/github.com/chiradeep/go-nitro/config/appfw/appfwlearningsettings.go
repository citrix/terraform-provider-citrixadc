package appfw

type Appfwlearningsettings struct {
	Contenttypeminthreshold                 int    `json:"contenttypeminthreshold,omitempty"`
	Contenttypepercentthreshold             int    `json:"contenttypepercentthreshold,omitempty"`
	Cookieconsistencyminthreshold           int    `json:"cookieconsistencyminthreshold,omitempty"`
	Cookieconsistencypercentthreshold       int    `json:"cookieconsistencypercentthreshold,omitempty"`
	Creditcardnumberminthreshold            int    `json:"creditcardnumberminthreshold,omitempty"`
	Creditcardnumberpercentthreshold        int    `json:"creditcardnumberpercentthreshold,omitempty"`
	Crosssitescriptingautodeploygraceperiod int    `json:"crosssitescriptingautodeploygraceperiod,omitempty"`
	Crosssitescriptingminthreshold          int    `json:"crosssitescriptingminthreshold,omitempty"`
	Crosssitescriptingpercentthreshold      int    `json:"crosssitescriptingpercentthreshold,omitempty"`
	Csrftagminthreshold                     int    `json:"csrftagminthreshold,omitempty"`
	Csrftagpercentthreshold                 int    `json:"csrftagpercentthreshold,omitempty"`
	Fieldconsistencyminthreshold            int    `json:"fieldconsistencyminthreshold,omitempty"`
	Fieldconsistencypercentthreshold        int    `json:"fieldconsistencypercentthreshold,omitempty"`
	Fieldformatautodeploygraceperiod        int    `json:"fieldformatautodeploygraceperiod,omitempty"`
	Fieldformatminthreshold                 int    `json:"fieldformatminthreshold,omitempty"`
	Fieldformatpercentthreshold             int    `json:"fieldformatpercentthreshold,omitempty"`
	Profilename                             string `json:"profilename,omitempty"`
	Sqlinjectionautodeploygraceperiod       int    `json:"sqlinjectionautodeploygraceperiod,omitempty"`
	Sqlinjectionminthreshold                int    `json:"sqlinjectionminthreshold,omitempty"`
	Sqlinjectionpercentthreshold            int    `json:"sqlinjectionpercentthreshold,omitempty"`
	Starturlautodeploygraceperiod           int    `json:"starturlautodeploygraceperiod,omitempty"`
	Starturlminthreshold                    int    `json:"starturlminthreshold,omitempty"`
	Starturlpercentthreshold                int    `json:"starturlpercentthreshold,omitempty"`
	Xmlattachmentminthreshold               int    `json:"xmlattachmentminthreshold,omitempty"`
	Xmlattachmentpercentthreshold           int    `json:"xmlattachmentpercentthreshold,omitempty"`
	Xmlwsiminthreshold                      int    `json:"xmlwsiminthreshold,omitempty"`
	Xmlwsipercentthreshold                  int    `json:"xmlwsipercentthreshold,omitempty"`
}
