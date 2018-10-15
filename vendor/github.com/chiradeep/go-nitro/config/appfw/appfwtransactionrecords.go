package appfw

type Appfwtransactionrecords struct {
	Appfwsessionid            string `json:"appfwsessionid,omitempty"`
	Clientip                  string `json:"clientip,omitempty"`
	Destip                    string `json:"destip,omitempty"`
	Endtime                   string `json:"endtime,omitempty"`
	Httptransactionid         int    `json:"httptransactionid,omitempty"`
	Nodeid                    int    `json:"nodeid,omitempty"`
	Packetengineid            int    `json:"packetengineid,omitempty"`
	Profilename               string `json:"profilename,omitempty"`
	Requestcontentlength      int    `json:"requestcontentlength,omitempty"`
	Requestmaxprocessingtime  int    `json:"requestmaxprocessingtime,omitempty"`
	Requestyields             int    `json:"requestyields,omitempty"`
	Responsecontentlength     int    `json:"responsecontentlength,omitempty"`
	Responsemaxprocessingtime int    `json:"responsemaxprocessingtime,omitempty"`
	Responseyields            int    `json:"responseyields,omitempty"`
	Starttime                 string `json:"starttime,omitempty"`
	Url                       string `json:"url,omitempty"`
}
