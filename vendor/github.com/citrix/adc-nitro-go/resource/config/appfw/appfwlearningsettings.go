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
	Starturlminthreshold uint32 `json:"starturlminthreshold,omitempty"`
	/**
	* Minimum percentage of application firewall sessions that must contain a particular start URL pattern for the learning engine to learn that start URL.
	*/
	Starturlpercentthreshold uint32 `json:"starturlpercentthreshold,omitempty"`
	/**
	* Minimum number of application firewall sessions that the learning engine must observe to learn cookies.
	*/
	Cookieconsistencyminthreshold uint32 `json:"cookieconsistencyminthreshold,omitempty"`
	/**
	* Minimum percentage of application firewall sessions that must contain a particular cookie pattern for the learning engine to learn that cookie.
	*/
	Cookieconsistencypercentthreshold uint32 `json:"cookieconsistencypercentthreshold,omitempty"`
	/**
	* Minimum number of application firewall sessions that the learning engine must observe to learn cross-site request forgery (CSRF) tags.
	*/
	Csrftagminthreshold uint32 `json:"csrftagminthreshold,omitempty"`
	/**
	* Minimum percentage of application firewall sessions that must contain a particular CSRF tag for the learning engine to learn that CSRF tag.
	*/
	Csrftagpercentthreshold uint32 `json:"csrftagpercentthreshold,omitempty"`
	/**
	* Minimum number of application firewall sessions that the learning engine must observe to learn field consistency information.
	*/
	Fieldconsistencyminthreshold uint32 `json:"fieldconsistencyminthreshold,omitempty"`
	/**
	* Minimum percentage of application firewall sessions that must contain a particular field consistency pattern for the learning engine to learn that field consistency pattern.
	*/
	Fieldconsistencypercentthreshold uint32 `json:"fieldconsistencypercentthreshold,omitempty"`
	/**
	* Minimum number of application firewall sessions that the learning engine must observe to learn HTML cross-site scripting patterns.
	*/
	Crosssitescriptingminthreshold uint32 `json:"crosssitescriptingminthreshold,omitempty"`
	/**
	* Minimum percentage of application firewall sessions that must contain a particular cross-site scripting pattern for the learning engine to learn that cross-site scripting pattern.
	*/
	Crosssitescriptingpercentthreshold uint32 `json:"crosssitescriptingpercentthreshold,omitempty"`
	/**
	* Minimum number of application firewall sessions that the learning engine must observe to learn HTML SQL injection patterns.
	*/
	Sqlinjectionminthreshold uint32 `json:"sqlinjectionminthreshold,omitempty"`
	/**
	* Minimum percentage of application firewall sessions that must contain a particular HTML SQL injection pattern for the learning engine to learn that HTML SQL injection pattern.
	*/
	Sqlinjectionpercentthreshold uint32 `json:"sqlinjectionpercentthreshold,omitempty"`
	/**
	* Minimum number of application firewall sessions that the learning engine must observe to learn field formats.
	*/
	Fieldformatminthreshold uint32 `json:"fieldformatminthreshold,omitempty"`
	/**
	* Minimum percentage of application firewall sessions that must contain a particular web form field pattern for the learning engine to recommend a field format for that form field.
	*/
	Fieldformatpercentthreshold uint32 `json:"fieldformatpercentthreshold,omitempty"`
	/**
	* Minimum threshold to learn Credit Card information.
	*/
	Creditcardnumberminthreshold uint32 `json:"creditcardnumberminthreshold,omitempty"`
	/**
	* Minimum threshold in percent to learn Credit Card information.
	*/
	Creditcardnumberpercentthreshold uint32 `json:"creditcardnumberpercentthreshold,omitempty"`
	/**
	* Minimum threshold to learn Content Type information.
	*/
	Contenttypeminthreshold uint32 `json:"contenttypeminthreshold,omitempty"`
	/**
	* Minimum threshold in percent to learn Content Type information.
	*/
	Contenttypepercentthreshold uint32 `json:"contenttypepercentthreshold,omitempty"`
	/**
	* Minimum number of application firewall sessions that the learning engine must observe to learn web services interoperability (WSI) information.
	*/
	Xmlwsiminthreshold uint32 `json:"xmlwsiminthreshold,omitempty"`
	/**
	* Minimum percentage of application firewall sessions that must contain a particular pattern for the learning engine to learn a web services interoperability (WSI) pattern.
	*/
	Xmlwsipercentthreshold uint32 `json:"xmlwsipercentthreshold,omitempty"`
	/**
	* Minimum number of application firewall sessions that the learning engine must observe to learn XML attachment patterns.
	*/
	Xmlattachmentminthreshold uint32 `json:"xmlattachmentminthreshold,omitempty"`
	/**
	* Minimum percentage of application firewall sessions that must contain a particular XML attachment pattern for the learning engine to learn that XML attachment pattern.
	*/
	Xmlattachmentpercentthreshold uint32 `json:"xmlattachmentpercentthreshold,omitempty"`
	/**
	* The number of minutes after the threshold hit alert the learned rule will be deployed
	*/
	Fieldformatautodeploygraceperiod uint32 `json:"fieldformatautodeploygraceperiod,omitempty"`
	/**
	* The number of minutes after the threshold hit alert the learned rule will be deployed
	*/
	Sqlinjectionautodeploygraceperiod uint32 `json:"sqlinjectionautodeploygraceperiod,omitempty"`
	/**
	* The number of minutes after the threshold hit alert the learned rule will be deployed
	*/
	Crosssitescriptingautodeploygraceperiod uint32 `json:"crosssitescriptingautodeploygraceperiod,omitempty"`
	/**
	* The number of minutes after the threshold hit alert the learned rule will be deployed
	*/
	Starturlautodeploygraceperiod uint32 `json:"starturlautodeploygraceperiod,omitempty"`
	/**
	* The number of minutes after the threshold hit alert the learned rule will be deployed
	*/
	Cookieconsistencyautodeploygraceperiod uint32 `json:"cookieconsistencyautodeploygraceperiod,omitempty"`
	/**
	* The number of minutes after the threshold hit alert the learned rule will be deployed
	*/
	Csrftagautodeploygraceperiod uint32 `json:"csrftagautodeploygraceperiod,omitempty"`
	/**
	* The number of minutes after the threshold hit alert the learned rule will be deployed
	*/
	Fieldconsistencyautodeploygraceperiod uint32 `json:"fieldconsistencyautodeploygraceperiod,omitempty"`
	/**
	* The number of minutes after the threshold hit alert the learned rule will be deployed
	*/
	Contenttypeautodeploygraceperiod uint32 `json:"contenttypeautodeploygraceperiod,omitempty"`

}
