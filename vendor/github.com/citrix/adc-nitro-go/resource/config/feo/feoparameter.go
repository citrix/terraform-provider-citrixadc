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

package feo

/**
* Configuration for FEO parameter resource.
*/
type Feoparameter struct {
	/**
	* The percentage value of a JPEG image quality to be reduced. Range: 0 - 100
	*/
	Jpegqualitypercent int `json:"jpegqualitypercent"`
	/**
	* Threshold value of the file size (in bytes) for converting external CSS files to inline CSS files.
	*/
	Cssinlinethressize int `json:"cssinlinethressize,omitempty"`
	/**
	* Threshold value of the file size (in bytes), for converting external JavaScript files to inline JavaScript files.
	*/
	Jsinlinethressize int `json:"jsinlinethressize,omitempty"`
	/**
	* Maximum file size of an image (in bytes), for coverting linked images to inline images.
	*/
	Imginlinethressize int `json:"imginlinethressize,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
