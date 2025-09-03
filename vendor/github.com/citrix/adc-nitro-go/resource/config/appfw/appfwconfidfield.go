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
* Configuration for configured confidential form fields resource.
*/
type Appfwconfidfield struct {
	/**
	* Name of the form field to designate as confidential.
	*/
	Fieldname string `json:"fieldname,omitempty"`
	/**
	* URL of the web page that contains the web form.
	*/
	Url string `json:"url,omitempty"`
	/**
	* Method of specifying the form field name. Available settings function as follows:
		* REGEX. Form field is a regular expression.
		* NOTREGEX. Form field is a literal string.
	*/
	Isregex string `json:"isregex,omitempty"`
	/**
	* Any comments to preserve information about the form field designation.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Enable or disable the confidential field designation.
	*/
	State string `json:"state,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
