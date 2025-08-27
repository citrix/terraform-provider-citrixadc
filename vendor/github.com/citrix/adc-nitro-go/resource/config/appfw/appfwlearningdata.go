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
* Configuration for learning data resource.
*/
type Appfwlearningdata struct {
	/**
	* Name of the profile.
	*/
	Profilename string `json:"profilename,omitempty"`
	/**
	* Start URL configuration.
	*/
	Starturl string `json:"starturl,omitempty"`
	/**
	* Cookie Name.
	*/
	Cookieconsistency string `json:"cookieconsistency,omitempty"`
	/**
	* Form field name.
	*/
	Fieldconsistency string `json:"fieldconsistency,omitempty"`
	/**
	* Form action URL.
	*/
	Formactionurlffc string `json:"formactionurl_ffc,omitempty"`
	/**
	* Content Type Name.
	*/
	Contenttype string `json:"contenttype,omitempty"`
	/**
	* Cross-site scripting.
	*/
	Crosssitescripting string `json:"crosssitescripting,omitempty"`
	/**
	* Form action URL.
	*/
	Formactionurlxss string `json:"formactionurl_xss,omitempty"`
	/**
	* Location of cross-site scripting exception - form field, header, cookie or url.
	*/
	Asscanlocationxss string `json:"as_scan_location_xss,omitempty"`
	/**
	* XSS value type. (Tag | Attribute | Pattern)
	*/
	Asvaluetypexss string `json:"as_value_type_xss,omitempty"`
	/**
	* XSS value expressions consistituting expressions for Tag, Attribute or Pattern.
	*/
	Asvalueexprxss string `json:"as_value_expr_xss,omitempty"`
	/**
	* Form field name.
	*/
	Sqlinjection string `json:"sqlinjection,omitempty"`
	/**
	* Form action URL.
	*/
	Formactionurlsql string `json:"formactionurl_sql,omitempty"`
	/**
	* Location of sql injection exception - form field, header or cookie.
	*/
	Asscanlocationsql string `json:"as_scan_location_sql,omitempty"`
	/**
	* SQL value type. Keyword, SpecialString or Wildchar
	*/
	Asvaluetypesql string `json:"as_value_type_sql,omitempty"`
	/**
	* SQL value expressions consistituting expressions for Keyword, SpecialString or Wildchar.
	*/
	Asvalueexprsql string `json:"as_value_expr_sql,omitempty"`
	/**
	* Field format name.
	*/
	Fieldformat string `json:"fieldformat,omitempty"`
	/**
	* Form action URL.
	*/
	Formactionurlff string `json:"formactionurl_ff,omitempty"`
	/**
	* CSRF Form Action URL
	*/
	Csrftag string `json:"csrftag,omitempty"`
	/**
	* CSRF Form Origin URL.
	*/
	Csrfformoriginurl string `json:"csrfformoriginurl,omitempty"`
	/**
	* The object expression that is to be excluded from safe commerce check.
	*/
	Creditcardnumber string `json:"creditcardnumber,omitempty"`
	/**
	* The url for which the list of credit card numbers are needed to be bypassed from inspection
	*/
	Creditcardnumberurl string `json:"creditcardnumberurl,omitempty"`
	/**
	* XML Denial of Service check, one of
		MaxAttributes
		MaxAttributeNameLength
		MaxAttributeValueLength
		MaxElementNameLength
		MaxFileSize
		MinFileSize
		MaxCDATALength
		MaxElements
		MaxElementDepth
		MaxElementChildren
		NumDTDs
		NumProcessingInstructions
		NumExternalEntities
		MaxEntityExpansions
		MaxEntityExpansionDepth
		MaxNamespaces
		MaxNamespaceUriLength
		MaxSOAPArraySize
		MaxSOAPArrayRank
	*/
	Xmldoscheck string `json:"xmldoscheck,omitempty"`
	/**
	* Web Services Interoperability Rule ID.
	*/
	Xmlwsicheck string `json:"xmlwsicheck,omitempty"`
	/**
	* XML Attachment Content-Type.
	*/
	Xmlattachmentcheck string `json:"xmlattachmentcheck,omitempty"`
	/**
	* Total XML requests.
	*/
	Totalxmlrequests bool `json:"totalxmlrequests,omitempty"`
	/**
	* Name of the security check.
	*/
	Securitycheck string `json:"securitycheck,omitempty"`
	/**
	* Target filename for data to be exported.
	*/
	Target string `json:"target,omitempty"`

	//------- Read only Parameter ---------;

	Url string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
	Fieldtype string `json:"fieldtype,omitempty"`
	Fieldformatminlength string `json:"fieldformatminlength,omitempty"`
	Fieldformatmaxlength string `json:"fieldformatmaxlength,omitempty"`
	Fieldformatcharmappcre string `json:"fieldformatcharmappcre,omitempty"`
	Valuetype string `json:"value_type,omitempty"`
	Value string `json:"value,omitempty"`
	Hits string `json:"hits,omitempty"`
	Data string `json:"data,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
