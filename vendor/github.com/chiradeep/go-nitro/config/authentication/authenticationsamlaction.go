package authentication

type Authenticationsamlaction struct {
	Artifactresolutionserviceurl   string      `json:"artifactresolutionserviceurl,omitempty"`
	Attribute1                     string      `json:"attribute1,omitempty"`
	Attribute10                    string      `json:"attribute10,omitempty"`
	Attribute11                    string      `json:"attribute11,omitempty"`
	Attribute12                    string      `json:"attribute12,omitempty"`
	Attribute13                    string      `json:"attribute13,omitempty"`
	Attribute14                    string      `json:"attribute14,omitempty"`
	Attribute15                    string      `json:"attribute15,omitempty"`
	Attribute16                    string      `json:"attribute16,omitempty"`
	Attribute2                     string      `json:"attribute2,omitempty"`
	Attribute3                     string      `json:"attribute3,omitempty"`
	Attribute4                     string      `json:"attribute4,omitempty"`
	Attribute5                     string      `json:"attribute5,omitempty"`
	Attribute6                     string      `json:"attribute6,omitempty"`
	Attribute7                     string      `json:"attribute7,omitempty"`
	Attribute8                     string      `json:"attribute8,omitempty"`
	Attribute9                     string      `json:"attribute9,omitempty"`
	Attributeconsumingserviceindex int         `json:"attributeconsumingserviceindex,omitempty"`
	Audience                       string      `json:"audience,omitempty"`
	Authnctxclassref               interface{} `json:"authnctxclassref,omitempty"`
	Defaultauthenticationgroup     string      `json:"defaultauthenticationgroup,omitempty"`
	Digestmethod                   string      `json:"digestmethod,omitempty"`
	Enforceusername                string      `json:"enforceusername,omitempty"`
	Forceauthn                     string      `json:"forceauthn,omitempty"`
	Groupnamefield                 string      `json:"groupnamefield,omitempty"`
	Logoutbinding                  string      `json:"logoutbinding,omitempty"`
	Logouturl                      string      `json:"logouturl,omitempty"`
	Metadataimportstatus           string      `json:"metadataimportstatus,omitempty"`
	Metadatarefreshinterval        int         `json:"metadatarefreshinterval,omitempty"`
	Metadataurl                    string      `json:"metadataurl,omitempty"`
	Name                           string      `json:"name,omitempty"`
	Requestedauthncontext          string      `json:"requestedauthncontext,omitempty"`
	Samlacsindex                   int         `json:"samlacsindex,omitempty"`
	Samlbinding                    string      `json:"samlbinding,omitempty"`
	Samlidpcertname                string      `json:"samlidpcertname,omitempty"`
	Samlissuername                 string      `json:"samlissuername,omitempty"`
	Samlredirecturl                string      `json:"samlredirecturl,omitempty"`
	Samlrejectunsignedassertion    string      `json:"samlrejectunsignedassertion,omitempty"`
	Samlsigningcertname            string      `json:"samlsigningcertname,omitempty"`
	Samltwofactor                  string      `json:"samltwofactor,omitempty"`
	Samluserfield                  string      `json:"samluserfield,omitempty"`
	Sendthumbprint                 string      `json:"sendthumbprint,omitempty"`
	Signaturealg                   string      `json:"signaturealg,omitempty"`
	Skewtime                       int         `json:"skewtime,omitempty"`
}
