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
* Configuration for feature resource.
*/
type Nsfeature struct {
	/**
	* Feature to be enabled. Multiple features can be specified by providing a blank space between each feature.
	*/
	Feature []string `json:"feature,omitempty"`

	//------- Read only Parameter ---------;

	Wl string `json:"wl,omitempty"`
	Sp string `json:"sp,omitempty"`
	Lb string `json:"lb,omitempty"`
	Cs string `json:"cs,omitempty"`
	Cr string `json:"cr,omitempty"`
	Sc string `json:"sc,omitempty"`
	Cmp string `json:"cmp,omitempty"`
	Pq string `json:"pq,omitempty"`
	Ssl string `json:"ssl,omitempty"`
	Gslb string `json:"gslb,omitempty"`
	Hdosp string `json:"hdosp,omitempty"`
	Cf string `json:"cf,omitempty"`
	Ic string `json:"ic,omitempty"`
	Sslvpn string `json:"sslvpn,omitempty"`
	Aaa string `json:"aaa,omitempty"`
	Ospf string `json:"ospf,omitempty"`
	Rip string `json:"rip,omitempty"`
	Bgp string `json:"bgp,omitempty"`
	Rewrite string `json:"rewrite,omitempty"`
	Ipv6pt string `json:"ipv6pt,omitempty"`
	Appfw string `json:"appfw,omitempty"`
	Responder string `json:"responder,omitempty"`
	Htmlinjection string `json:"htmlinjection,omitempty"`
	Push string `json:"push,omitempty"`
	Appflow string `json:"appflow,omitempty"`
	Cloudbridge string `json:"cloudbridge,omitempty"`
	Isis string `json:"isis,omitempty"`
	Ch string `json:"ch,omitempty"`
	Appqoe string `json:"appqoe,omitempty"`
	Contentaccelerator string `json:"contentaccelerator,omitempty"`
	Feo string `json:"feo,omitempty"`
	Lsn string `json:"lsn,omitempty"`
	Rdpproxy string `json:"rdpproxy,omitempty"`
	Rep string `json:"rep,omitempty"`
	Urlfiltering string `json:"urlfiltering,omitempty"`
	Videooptimization string `json:"videooptimization,omitempty"`
	Forwardproxy string `json:"forwardproxy,omitempty"`
	Sslinterception string `json:"sslinterception,omitempty"`
	Adaptivetcp string `json:"adaptivetcp,omitempty"`
	Cqa string `json:"cqa,omitempty"`
	Ci string `json:"ci,omitempty"`
	Bot string `json:"bot,omitempty"`
	Apigateway string `json:"apigateway,omitempty"`

}
