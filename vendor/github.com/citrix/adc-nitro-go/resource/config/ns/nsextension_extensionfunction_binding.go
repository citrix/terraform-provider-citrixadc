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

package ns

/**
* Binding class showing the extensionfunction that can be bound to nsextension.
*/
type Nsextensionextensionfunctionbinding struct {
	/**
	* Name of extension function given in the extension.
	*/
	Extensionfunctionname string `json:"extensionfunctionname,omitempty"`
	/**
	* Line number of the function in file.
	*/
	Extensionfunctionlinenumber int `json:"extensionfunctionlinenumber,omitempty"`
	/**
	* Extension function class type.
	*/
	Extensionfunctionclasstype string `json:"extensionfunctionclasstype,omitempty"`
	/**
	* Extension function return type.
	*/
	Extensionfunctionreturntype string `json:"extensionfunctionreturntype,omitempty"`
	/**
	* Extension function is in use or not.
	*/
	Activeextensionfunction int `json:"activeextensionfunction,omitempty"`
	/**
	* List of extension function's arguments types
	*/
	Extensionfunctionargtype []string `json:"extensionfunctionargtype,omitempty"`
	/**
	* Any description to preserve information about the extension function.
	*/
	Extensionfuncdescription string `json:"extensionfuncdescription,omitempty"`
	/**
	* Number of parameters in the extension function
	*/
	Extensionfunctionargcount int `json:"extensionfunctionargcount,omitempty"`
	/**
	* List of classes (including inherited) that the function is present in.
	*/
	Extensionfunctionclasses []string `json:"extensionfunctionclasses,omitempty"`
	/**
	* Number of classes the function is present in.
	*/
	Extensionfunctionclassescount int `json:"extensionfunctionclassescount,omitempty"`
	/**
	* List of parameters (including promotions) that the function can accept.
	*/
	Extensionfunctionallparams []string `json:"extensionfunctionallparams,omitempty"`
	/**
	* Number of parameters (including promotions) that the function can accept.
	*/
	Extensionfunctionallparamscount int `json:"extensionfunctionallparamscount,omitempty"`
	/**
	* Name of the extension object.
	*/
	Name string `json:"name,omitempty"`


}