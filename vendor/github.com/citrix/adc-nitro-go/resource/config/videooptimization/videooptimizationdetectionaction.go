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

package videooptimization

/**
* Configuration for videooptimization detectionaction resource.
*/
type Videooptimizationdetectionaction struct {
	/**
	* Name for the video optimization detection action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Type of video optimization action. Available settings function as follows:
		* clear_text_pd - Cleartext PD type is detected.
		* clear_text_abr - Cleartext ABR is detected.
		* encrypted_abr - Encrypted ABR is detected.
		* trigger_enc_abr - Possible encrypted ABR is detected.
		* trigger_body_detection - Possible cleartext ABR is detected. Triggers body content detection.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Comment. Any type of information about this video optimization detection action.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* New name for the videooptimization detection action.
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Referencecount string `json:"referencecount,omitempty"`
	Undefhits string `json:"undefhits,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
