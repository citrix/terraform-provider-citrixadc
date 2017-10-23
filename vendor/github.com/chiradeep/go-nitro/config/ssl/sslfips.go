package ssl

type Sslfips struct {
	Coresenabled        int    `json:"coresenabled,omitempty"`
	Coresmax            int    `json:"coresmax,omitempty"`
	Erasedata           string `json:"erasedata,omitempty"`
	Fipsfw              string `json:"fipsfw,omitempty"`
	Firmwarereleasedate string `json:"firmwarereleasedate,omitempty"`
	Flag                int    `json:"flag,omitempty"`
	Flashmemoryfree     int    `json:"flashmemoryfree,omitempty"`
	Flashmemorytotal    int    `json:"flashmemorytotal,omitempty"`
	Hsmlabel            string `json:"hsmlabel,omitempty"`
	Inithsm             string `json:"inithsm,omitempty"`
	Majorversion        int    `json:"majorversion,omitempty"`
	Minorversion        int    `json:"minorversion,omitempty"`
	Oldsopassword       string `json:"oldsopassword,omitempty"`
	Serial              int    `json:"serial,omitempty"`
	Serialno            string `json:"serialno,omitempty"`
	Sopassword          string `json:"sopassword,omitempty"`
	Sramfree            int    `json:"sramfree,omitempty"`
	Sramtotal           int    `json:"sramtotal,omitempty"`
	State               int    `json:"state,omitempty"`
	Status              int    `json:"status,omitempty"`
	Userpassword        string `json:"userpassword,omitempty"`
}
