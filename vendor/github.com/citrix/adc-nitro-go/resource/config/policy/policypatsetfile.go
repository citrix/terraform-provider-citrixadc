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

package policy

/**
* Configuration for patset file resource.
*/
type Policypatsetfile struct {
	/**
	* URL in protocol, host, path, and file name format from where the patset file will be imported. If file is already present, then it can be imported using local keyword (import patsetfile local:filename patsetfile1)
		NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access
	*/
	Src string `json:"src,omitempty"`
	/**
	* Name to assign to the imported patset file. Unique name of the pattern set. Not case sensitive. Must begin with an ASCII letter or underscore (_) character and must contain only alphanumeric and underscore characters.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Overwrites the existing file
	*/
	Overwrite bool `json:"overwrite,omitempty"`
	/**
	* patset file patterns delimiter.
	*/
	Delimiter string `json:"delimiter,omitempty"`
	/**
	* Character set associated with the characters in the string.
	*/
	Charset string `json:"charset,omitempty"`
	/**
	* Any comments to preserve information about this patsetfile.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* When set, display only shows all imported patsetfiles.
	*/
	Imported bool `json:"imported,omitempty"`

	//------- Read only Parameter ---------;

	Totalpatterns string `json:"totalpatterns,omitempty"`
	Boundpatterns string `json:"boundpatterns,omitempty"`
	Patsetname string `json:"patsetname,omitempty"`
	Bindstatuscode string `json:"bindstatuscode,omitempty"`
	Bindstatus string `json:"bindstatus,omitempty"`

}
