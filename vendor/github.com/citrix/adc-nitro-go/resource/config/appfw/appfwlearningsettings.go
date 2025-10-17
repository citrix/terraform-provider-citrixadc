/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package appfw

/**
* Configuration for learning settings resource.
*/
type Appfwlearningsettings struct {
	/**
	* Name of the profile.
	*/
	Profilename string `json:"profilename,omitempty"`
	/**
	* Minimum number of application firewall sessions that the learning engine must observe to learn start URLs.
	*/
	Starturlminthreshold *int `json:"starturlminthreshold,omitempty"`
	/**
	* Minimum percentage of application firewall sessions that must contain a particular start URL pattern for the learning engine to learn that start URL.
	*/
	Starturlpercentthreshold *int `json:"starturlpercentthreshold,omitempty"`
	/**
	* Minimum number of application firewall sessions that the learning engine must observe to learn cookies.
	*/
	Cookieconsistencyminthreshold *int `json:"cookieconsistencyminthreshold,omitempty"`
	/**
	* Minimum percentage of application firewall sessions that must contain a particular cookie pattern for the learning engine to learn that cookie.
	*/
	Cookieconsistencypercentthreshold *int `json:"cookieconsistencypercentthreshold,omitempty"`
	/**
	* Minimum number of application firewall sessions that the learning engine must observe to learn cross-site request forgery (CSRF) tags.
	*/
	Csrftagminthreshold *int `json:"csrftagminthreshold,omitempty"`
	/**
	* Minimum percentage of application firewall sessions that must contain a particular CSRF tag for the learning engine to learn that CSRF tag.
	*/
	Csrftagpercentthreshold *int `json:"csrftagpercentthreshold,omitempty"`
	/**
	* Minimum number of application firewall sessions that the learning engine must observe to learn field consistency information.
	*/
	Fieldconsistencyminthreshold *int `json:"fieldconsistencyminthreshold,omitempty"`
	/**
	* Minimum percentage of application firewall sessions that must contain a particular field consistency pattern for the learning engine to learn that field consistency pattern.
	*/
	Fieldconsistencypercentthreshold *int `json:"fieldconsistencypercentthreshold,omitempty"`
	/**
	* Minimum number of application firewall sessions that the learning engine must observe to learn HTML cross-site scripting patterns.
	*/
	Crosssitescriptingminthreshold *int `json:"crosssitescriptingminthreshold,omitempty"`
	/**
	* Minimum percentage of application firewall sessions that must contain a particular cross-site scripting pattern for the learning engine to learn that cross-site scripting pattern.
	*/
	Crosssitescriptingpercentthreshold *int `json:"crosssitescriptingpercentthreshold,omitempty"`
	/**
	* Minimum number of application firewall sessions that the learning engine must observe to learn HTML SQL injection patterns.
	*/
	Sqlinjectionminthreshold *int `json:"sqlinjectionminthreshold,omitempty"`
	/**
	* Minimum percentage of application firewall sessions that must contain a particular HTML SQL injection pattern for the learning engine to learn that HTML SQL injection pattern.
	*/
	Sqlinjectionpercentthreshold *int `json:"sqlinjectionpercentthreshold,omitempty"`
	/**
	* Minimum number of application firewall sessions that the learning engine must observe to learn field formats.
	*/
	Fieldformatminthreshold *int `json:"fieldformatminthreshold,omitempty"`
	/**
	* Minimum percentage of application firewall sessions that must contain a particular web form field pattern for the learning engine to recommend a field format for that form field.
	*/
	Fieldformatpercentthreshold *int `json:"fieldformatpercentthreshold,omitempty"`
	/**
	* Minimum threshold to learn Credit Card information.
	*/
	Creditcardnumberminthreshold *int `json:"creditcardnumberminthreshold,omitempty"`
	/**
	* Minimum threshold in percent to learn Credit Card information.
	*/
	Creditcardnumberpercentthreshold *int `json:"creditcardnumberpercentthreshold,omitempty"`
	/**
	* Minimum threshold to learn Content Type information.
	*/
	Contenttypeminthreshold *int `json:"contenttypeminthreshold,omitempty"`
	/**
	* Minimum threshold in percent to learn Content Type information.
	*/
	Contenttypepercentthreshold *int `json:"contenttypepercentthreshold,omitempty"`
	/**
	* Minimum number of application firewall sessions that the learning engine must observe to learn web services interoperability (WSI) information.
	*/
	Xmlwsiminthreshold *int `json:"xmlwsiminthreshold,omitempty"`
	/**
	* Minimum percentage of application firewall sessions that must contain a particular pattern for the learning engine to learn a web services interoperability (WSI) pattern.
	*/
	Xmlwsipercentthreshold *int `json:"xmlwsipercentthreshold,omitempty"`
	/**
	* Minimum number of application firewall sessions that the learning engine must observe to learn XML attachment patterns.
	*/
	Xmlattachmentminthreshold *int `json:"xmlattachmentminthreshold,omitempty"`
	/**
	* Minimum percentage of application firewall sessions that must contain a particular XML attachment pattern for the learning engine to learn that XML attachment pattern.
	*/
	Xmlattachmentpercentthreshold *int `json:"xmlattachmentpercentthreshold,omitempty"`
	/**
	* The number of minutes after the threshold hit alert the learned rule will be deployed
	*/
	Fieldformatautodeploygraceperiod *int `json:"fieldformatautodeploygraceperiod,omitempty"`
	/**
	* The number of minutes after the threshold hit alert the learned rule will be deployed
	*/
	Sqlinjectionautodeploygraceperiod *int `json:"sqlinjectionautodeploygraceperiod,omitempty"`
	/**
	* The number of minutes after the threshold hit alert the learned rule will be deployed
	*/
	Crosssitescriptingautodeploygraceperiod *int `json:"crosssitescriptingautodeploygraceperiod,omitempty"`
	/**
	* The number of minutes after the threshold hit alert the learned rule will be deployed
	*/
	Starturlautodeploygraceperiod *int `json:"starturlautodeploygraceperiod,omitempty"`
	/**
	* The number of minutes after the threshold hit alert the learned rule will be deployed
	*/
	Cookieconsistencyautodeploygraceperiod *int `json:"cookieconsistencyautodeploygraceperiod,omitempty"`
	/**
	* The number of minutes after the threshold hit alert the learned rule will be deployed
	*/
	Csrftagautodeploygraceperiod *int `json:"csrftagautodeploygraceperiod,omitempty"`
	/**
	* The number of minutes after the threshold hit alert the learned rule will be deployed
	*/
	Fieldconsistencyautodeploygraceperiod *int `json:"fieldconsistencyautodeploygraceperiod,omitempty"`
	/**
	* The number of minutes after the threshold hit alert the learned rule will be deployed
	*/
	Contenttypeautodeploygraceperiod *int `json:"contenttypeautodeploygraceperiod,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
