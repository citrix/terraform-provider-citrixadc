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
* Configuration for Front end optimization action resource.
*/
type Feoaction struct {
	/**
	* The name of the front end optimization action.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Extend the time period during which the browser can use the cached resource.
	*/
	Pageextendcache bool `json:"pageextendcache,omitempty"`
	/**
	* Maxage for cache extension.
	*/
	Cachemaxage *int `json:"cachemaxage"` // Zero is a valid value
	/**
	* Shrink image dimensions as per the height and width attributes specified in the <img> tag.
	*/
	Imgshrinktoattrib bool `json:"imgshrinktoattrib,omitempty"`
	/**
	* Convert GIF image formats to PNG formats.
	*/
	Imggiftopng bool `json:"imggiftopng,omitempty"`
	/**
	* Convert JPEG, GIF, PNG image formats to WEBP format.
	*/
	Imgtowebp bool `json:"imgtowebp,omitempty"`
	/**
	* Convert JPEG, GIF, PNG image formats to JXR format.
	*/
	Imgtojpegxr bool `json:"imgtojpegxr,omitempty"`
	/**
	* Inline images whose size is less than 2KB.
	*/
	Imginline bool `json:"imginline,omitempty"`
	/**
	* Inline small images (less than 2KB) referred within CSS files as background-URLs
	*/
	Cssimginline bool `json:"cssimginline,omitempty"`
	/**
	* Remove non-image data such as comments from JPEG images.
	*/
	Jpgoptimize bool `json:"jpgoptimize,omitempty"`
	/**
	* Download images, only when the user scrolls the page to view them.
	*/
	Imglazyload bool `json:"imglazyload,omitempty"`
	/**
	* Remove comments and whitespaces from CSSs.
	*/
	Cssminify bool `json:"cssminify,omitempty"`
	/**
	* Inline CSS files, whose size is less than 2KB, within the main page.
	*/
	Cssinline bool `json:"cssinline,omitempty"`
	/**
	* Combine one or more CSS files into one file.
	*/
	Csscombine bool `json:"csscombine,omitempty"`
	/**
	* Convert CSS import statements to HTML link tags.
	*/
	Convertimporttolink bool `json:"convertimporttolink,omitempty"`
	/**
	* Remove comments and whitespaces from JavaScript.
	*/
	Jsminify bool `json:"jsminify,omitempty"`
	/**
	* Convert linked JavaScript files (less than 2KB) to inline JavaScript files.
	*/
	Jsinline bool `json:"jsinline,omitempty"`
	/**
	* Remove comments and whitespaces from an HTML page.
	*/
	Htmlminify bool `json:"htmlminify,omitempty"`
	/**
	* Move any CSS file present within the body tag of an HTML page to the head tag.
	*/
	Cssmovetohead bool `json:"cssmovetohead,omitempty"`
	/**
	* Move any JavaScript present in the body tag to the end of the body tag.
	*/
	Jsmovetoend bool `json:"jsmovetoend,omitempty"`
	/**
	* Domain name of the server
	*/
	Domainsharding string `json:"domainsharding,omitempty"`
	/**
	* Set of domain names that replaces the parent domain.
	*/
	Dnsshards []string `json:"dnsshards,omitempty"`
	/**
	* Send AppFlow records about the web pages optimized by this action. The records provide FEO statistics, such as the number of HTTP requests that have been reduced for this page. You must enable the Appflow feature before enabling this parameter.
	*/
	Clientsidemeasurements bool `json:"clientsidemeasurements,omitempty"`

	//------- Read only Parameter ---------;

	Imgadddimensions string `json:"imgadddimensions,omitempty"`
	Imgshrinkformobile string `json:"imgshrinkformobile,omitempty"`
	Imgweaken string `json:"imgweaken,omitempty"`
	Jpgprogressive string `json:"jpgprogressive,omitempty"`
	Cssflattenimports string `json:"cssflattenimports,omitempty"`
	Jscombine string `json:"jscombine,omitempty"`
	Htmlrmdefaultattribs string `json:"htmlrmdefaultattribs,omitempty"`
	Htmlrmattribquotes string `json:"htmlrmattribquotes,omitempty"`
	Htmltrimurls string `json:"htmltrimurls,omitempty"`
	Hits string `json:"hits,omitempty"`
	Undefhits string `json:"undefhits,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
