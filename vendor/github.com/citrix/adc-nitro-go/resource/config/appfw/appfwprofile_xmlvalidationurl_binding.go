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
* Binding class showing the xmlvalidationurl that can be bound to appfwprofile.
*/
type Appfwprofilexmlvalidationurlbinding struct {
	/**
	* XML Validation URL regular expression.
	*/
	Xmlvalidationurl string `json:"xmlvalidationurl,omitempty"`
	/**
	* Validate response message.
	*/
	Xmlvalidateresponse string `json:"xmlvalidateresponse,omitempty"`
	/**
	* WSDL object for soap request validation.
	*/
	Xmlwsdl string `json:"xmlwsdl,omitempty"`
	/**
	* Allow addtional soap headers.
	*/
	Xmladditionalsoapheaders string `json:"xmladditionalsoapheaders,omitempty"`
	/**
	* Modifies the behaviour of the Request URL validation w.r.t. the Service URL.
		If set to ABSOLUTE, the entire request URL is validated with the entire URL mentioned in Service of the associated WSDL.
		eg: Service URL: http://example.org/ExampleService, Request URL: http//example.com/ExampleService would FAIL the validation.
		If set to RELAIVE, only the non-hostname part of the request URL is validated against the non-hostname part of the Service URL.
		eg: Service URL: http://example.org/ExampleService, Request URL: http//example.com/ExampleService would PASS the validation.
	*/
	Xmlendpointcheck string `json:"xmlendpointcheck,omitempty"`
	/**
	* XML Schema object for request validation .
	*/
	Xmlrequestschema string `json:"xmlrequestschema,omitempty"`
	/**
	* XML Schema object for response validation.
	*/
	Xmlresponseschema string `json:"xmlresponseschema,omitempty"`
	/**
	* Validate SOAP Evelope only.
	*/
	Xmlvalidatesoapenvelope string `json:"xmlvalidatesoapenvelope,omitempty"`
	/**
	* Enabled.
	*/
	State string `json:"state,omitempty"`
	/**
	* Any comments about the purpose of profile, or other useful information about the profile.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Is the rule auto deployed by dynamic profile ?
	*/
	Isautodeployed string `json:"isautodeployed,omitempty"`
	/**
	* Send SNMP alert?
	*/
	Alertonly string `json:"alertonly,omitempty"`
	/**
	* Name of the profile to which to bind an exemption or rule.
	*/
	Name string `json:"name,omitempty"`
	/**
	* A "id" that identifies the rule.
	*/
	Resourceid string `json:"resourceid,omitempty"`
	/**
	* Specifies rule type of binding
	*/
	Ruletype string `json:"ruletype,omitempty"`


}