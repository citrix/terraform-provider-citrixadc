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

package transform

/**
* Binding class showing the transformaction that can be bound to transformprofile.
*/
type Transformprofiletransformactionbinding struct {
	/**
	* URL Transformation action name.
	*/
	Actionname string `json:"actionname,omitempty"`
	/**
	* Priority of the Action within the Profile.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Enabled flag.
	*/
	State string `json:"state,omitempty"`
	/**
	* URL Transformation profile name.
	*/
	Profilename string `json:"profilename,omitempty"`
	/**
	* Pattern of original request URLs. It corresponds to the way external users view the server, and acts as a source for request transformations.
	*/
	Requrlfrom string `json:"requrlfrom,omitempty"`
	/**
	* Pattern of transformed request URLs. It corresponds to internal addresses and indicates how they are created.
	*/
	Requrlinto string `json:"requrlinto,omitempty"`
	/**
	* Pattern of original response URLs. It corresponds to the way external users view the server, and acts as a source for response transformations.
	*/
	Resurlfrom string `json:"resurlfrom,omitempty"`
	/**
	* Pattern of transformed response URLs. It corresponds to internal addresses and indicates how they are created.
	*/
	Resurlinto string `json:"resurlinto,omitempty"`
	/**
	* Pattern of the original domain in Set-Cookie headers.
	*/
	Cookiedomainfrom string `json:"cookiedomainfrom,omitempty"`
	/**
	* Pattern of the transformed domain in Set-Cookie headers.
	*/
	Cookiedomaininto string `json:"cookiedomaininto,omitempty"`
	/**
	* Comments.
	*/
	Actioncomment string `json:"actioncomment,omitempty"`
	/**
	* Name of the profile.
	*/
	Name string `json:"name,omitempty"`


}