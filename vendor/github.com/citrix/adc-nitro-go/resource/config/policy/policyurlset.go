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
* Configuration for URL set resource.
*/
type Policyurlset struct {
	/**
	* Unique name of the url set. Not case sensitive. Must begin with an ASCII letter or underscore (_) character and must contain only alphanumeric and underscore characters. Must not be the name of an existing named expression, pattern set, dataset, string map, or HTTP callout.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Any comments to preserve information about this url set.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* when set, display shows all imported urlsets.
	*/
	Imported bool `json:"imported,omitempty"`
	/**
	* Overwrites the existing file.
	*/
	Overwrite bool `json:"overwrite,omitempty"`
	/**
	* CSV file record delimiter.
	*/
	Delimiter string `json:"delimiter,omitempty"`
	/**
	* CSV file row separator.
	*/
	Rowseparator string `json:"rowseparator,omitempty"`
	/**
	* URL (protocol, host, path and file name) from where the CSV (comma separated file) file will be imported or exported. Each record/line will one entry within the urlset. The first field contains the URL pattern, subsequent fields contains the metadata, if available. HTTP, HTTPS and FTP protocols are supported. NOTE: The operation fails if the destination HTTPS server requires client certificate authentication for access.
	*/
	Url string `json:"url,omitempty"`
	/**
	* The interval, in seconds, rounded down to the nearest 15 minutes, at which the update of urlset occurs.
	*/
	Interval *int `json:"interval,omitempty"`
	/**
	* Prevent this urlset from being exported.
	*/
	Privateset bool `json:"privateset,omitempty"`
	/**
	* Force exact subdomain matching, ex. given an entry 'google.com' in the urlset, a request to 'news.google.com' won't match, if subdomainExactMatch is set.
	*/
	Subdomainexactmatch bool `json:"subdomainexactmatch,omitempty"`
	/**
	* An ID that would be sent to AppFlow to indicate which URLSet was the last one that matched the requested URL.
	*/
	Matchedid *int `json:"matchedid,omitempty"`
	/**
	* Add this URL to this urlset. Used for testing when contents of urlset is kept confidential.
	*/
	Canaryurl string `json:"canaryurl,omitempty"`

	//------- Read only Parameter ---------;

	Patterncount string `json:"patterncount,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
