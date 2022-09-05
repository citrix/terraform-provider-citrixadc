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

package analytics

/**
* Configuration for Analytics profile resource.
*/
type Analyticsprofile struct {
	/**
	* Name for the analytics profile. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at
		(@), equals (=), and hyphen (-) characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my appflow profile" or 'my appflow profile').
	*/
	Name string `json:"name,omitempty"`
	/**
	* The collector can be an IP, an appflow collector name, a service or a vserver. If IP is specified, the transport is considered as logstream and default port of 5557 is taken. If collector name is specified, the collector properties are taken from the configured collector. If service is specified, the configured service is assumed as the collector. If vserver is specified, the services bound to it are considered as collectors and the records are load balanced.
	*/
	Collectors string `json:"collectors,omitempty"`
	/**
	* This option indicates what information needs to be collected and exported.
	*/
	Type string `json:"type,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will insert a javascript into the HTTP response to collect the client side page-timings and will send the same to the configured collectors.
	*/
	Httpclientsidemeasurements string `json:"httpclientsidemeasurements,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will link the embedded objects of a page together.
	*/
	Httppagetracking string `json:"httppagetracking,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will log the URL in appflow records
	*/
	Httpurl string `json:"httpurl,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will log the Host header in appflow records
	*/
	Httphost string `json:"httphost,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will log the method header in appflow records
	*/
	Httpmethod string `json:"httpmethod,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will log the referer header in appflow records
	*/
	Httpreferer string `json:"httpreferer,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will log User-Agent header.
	*/
	Httpuseragent string `json:"httpuseragent,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will log cookie header.
	*/
	Httpcookie string `json:"httpcookie,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will log location header.
	*/
	Httplocation string `json:"httplocation,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will send the URL category record.
	*/
	Urlcategory string `json:"urlcategory,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will log all the request and response headers.
	*/
	Allhttpheaders string `json:"allhttpheaders,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will log content-length header.
	*/
	Httpcontenttype string `json:"httpcontenttype,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will log Authentication header.
	*/
	Httpauthentication string `json:"httpauthentication,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will Via header.
	*/
	Httpvia string `json:"httpvia,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will log X-Forwarded-For header.
	*/
	Httpxforwardedforheader string `json:"httpxforwardedforheader,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will log set-cookie header.
	*/
	Httpsetcookie string `json:"httpsetcookie,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will log set-cookie2 header.
	*/
	Httpsetcookie2 string `json:"httpsetcookie2,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will log domain name.
	*/
	Httpdomainname string `json:"httpdomainname,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will log URL Query.
	*/
	Httpurlquery string `json:"httpurlquery,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will log TCP burst parameters.
	*/
	Tcpburstreporting string `json:"tcpburstreporting,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will log TCP CQA parameters.
	*/
	Cqareporting string `json:"cqareporting,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will log the Integrated Caching appflow records
	*/
	Integratedcache string `json:"integratedcache,omitempty"`
	/**
	* On enabling this option, the Citrix ADC will log the gRPC status headers
	*/
	Grpcstatus string `json:"grpcstatus,omitempty"`
	/**
	* This option indicates the format of REST API POST body. It depends on the consumer of the analytics data.
	*/
	Outputmode string `json:"outputmode,omitempty"`
	/**
	* This option indicates the whether metrics should be sent to the REST collector.
	*/
	Metrics string `json:"metrics,omitempty"`
	/**
	* This option indicates the whether events should be sent to the REST collector.
	*/
	Events string `json:"events,omitempty"`
	/**
	* This option indicates the whether auditlog should be sent to the REST collector.
	*/
	Auditlogs string `json:"auditlogs,omitempty"`

	//------- Read only Parameter ---------;

	Refcnt string `json:"refcnt,omitempty"`

}
