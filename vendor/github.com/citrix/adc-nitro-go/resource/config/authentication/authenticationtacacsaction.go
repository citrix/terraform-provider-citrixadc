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

package authentication

/**
* Configuration for TACACS action resource.
*/
type Authenticationtacacsaction struct {
	/**
	* Name for the TACACS+ profile (action).
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after TACACS profile is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication action" or 'y authentication action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* IP address assigned to the TACACS+ server.
	*/
	Serverip string `json:"serverip,omitempty"`
	/**
	* Port number on which the TACACS+ server listens for connections.
	*/
	Serverport int `json:"serverport,omitempty"`
	/**
	* Number of seconds the Citrix ADC waits for a response from the TACACS+ server.
	*/
	Authtimeout int `json:"authtimeout,omitempty"`
	/**
	* Key shared between the TACACS+ server and the Citrix ADC.
		Required for allowing the Citrix ADC to communicate with the TACACS+ server.
	*/
	Tacacssecret string `json:"tacacssecret,omitempty"`
	/**
	* Use streaming authorization on the TACACS+ server.
	*/
	Authorization string `json:"authorization,omitempty"`
	/**
	* Whether the TACACS+ server is currently accepting accounting messages.
	*/
	Accounting string `json:"accounting,omitempty"`
	/**
	* The state of the TACACS+ server that will receive accounting messages.
	*/
	Auditfailedcmds string `json:"auditfailedcmds,omitempty"`
	/**
	* TACACS+ group attribute name.
		Used for group extraction on the TACACS+ server.
	*/
	Groupattrname string `json:"groupattrname,omitempty"`
	/**
	* This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
	*/
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	/**
	* Name of the custom attribute to be extracted from server and stored at index '1' (where '1' changes for each attribute)
	*/
	Attribute1 string `json:"attribute1,omitempty"`
	/**
	* Name of the custom attribute to be extracted from server and stored at index '2' (where '2' changes for each attribute)
	*/
	Attribute2 string `json:"attribute2,omitempty"`
	/**
	* Name of the custom attribute to be extracted from server and stored at index '3' (where '3' changes for each attribute)
	*/
	Attribute3 string `json:"attribute3,omitempty"`
	/**
	* Name of the custom attribute to be extracted from server and stored at index '4' (where '4' changes for each attribute)
	*/
	Attribute4 string `json:"attribute4,omitempty"`
	/**
	* Name of the custom attribute to be extracted from server and stored at index '5' (where '5' changes for each attribute)
	*/
	Attribute5 string `json:"attribute5,omitempty"`
	/**
	* Name of the custom attribute to be extracted from server and stored at index '6' (where '6' changes for each attribute)
	*/
	Attribute6 string `json:"attribute6,omitempty"`
	/**
	* Name of the custom attribute to be extracted from server and stored at index '7' (where '7' changes for each attribute)
	*/
	Attribute7 string `json:"attribute7,omitempty"`
	/**
	* Name of the custom attribute to be extracted from server and stored at index '8' (where '8' changes for each attribute)
	*/
	Attribute8 string `json:"attribute8,omitempty"`
	/**
	* Name of the custom attribute to be extracted from server and stored at index '9' (where '9' changes for each attribute)
	*/
	Attribute9 string `json:"attribute9,omitempty"`
	/**
	* Name of the custom attribute to be extracted from server and stored at index '10' (where '10' changes for each attribute)
	*/
	Attribute10 string `json:"attribute10,omitempty"`
	/**
	* Name of the custom attribute to be extracted from server and stored at index '11' (where '11' changes for each attribute)
	*/
	Attribute11 string `json:"attribute11,omitempty"`
	/**
	* Name of the custom attribute to be extracted from server and stored at index '12' (where '12' changes for each attribute)
	*/
	Attribute12 string `json:"attribute12,omitempty"`
	/**
	* Name of the custom attribute to be extracted from server and stored at index '13' (where '13' changes for each attribute)
	*/
	Attribute13 string `json:"attribute13,omitempty"`
	/**
	* Name of the custom attribute to be extracted from server and stored at index '14' (where '14' changes for each attribute)
	*/
	Attribute14 string `json:"attribute14,omitempty"`
	/**
	* Name of the custom attribute to be extracted from server and stored at index '15' (where '15' changes for each attribute)
	*/
	Attribute15 string `json:"attribute15,omitempty"`
	/**
	* Name of the custom attribute to be extracted from server and stored at index '16' (where '16' changes for each attribute)
	*/
	Attribute16 string `json:"attribute16,omitempty"`
	/**
	* List of attribute names separated by ',' which needs to be fetched from tacacs server.
		Note that preceeding and trailing spaces will be removed.
		Attribute name can be 127 bytes and total length of this string should not cross 2047 bytes.
		These attributes have multi-value support separated by ',' and stored as key-value pair in AAA session
	*/
	Attributes string `json:"attributes,omitempty"`

	//------- Read only Parameter ---------;

	Success string `json:"success,omitempty"`
	Failure string `json:"failure,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
