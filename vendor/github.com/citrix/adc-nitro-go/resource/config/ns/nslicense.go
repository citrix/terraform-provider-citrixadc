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
* Configuration for license resource.
*/
type Nslicense struct {

	//------- Read only Parameter ---------;

	Wl string `json:"wl,omitempty"`
	Sp string `json:"sp,omitempty"`
	Lb string `json:"lb,omitempty"`
	Cs string `json:"cs,omitempty"`
	Cr string `json:"cr,omitempty"`
	Cmp string `json:"cmp,omitempty"`
	Delta string `json:"delta,omitempty"`
	Ssl string `json:"ssl,omitempty"`
	Gslb string `json:"gslb,omitempty"`
	Gslbp string `json:"gslbp,omitempty"`
	Routing string `json:"routing,omitempty"`
	Cf string `json:"cf,omitempty"`
	Contentaccelerator string `json:"contentaccelerator,omitempty"`
	Ic string `json:"ic,omitempty"`
	Sslvpn string `json:"sslvpn,omitempty"`
	Fsslvpnusers string `json:"f_sslvpn_users,omitempty"`
	Ficausers string `json:"f_ica_users,omitempty"`
	Aaa string `json:"aaa,omitempty"`
	Ospf string `json:"ospf,omitempty"`
	Rip string `json:"rip,omitempty"`
	Bgp string `json:"bgp,omitempty"`
	Rewrite string `json:"rewrite,omitempty"`
	Ipv6pt string `json:"ipv6pt,omitempty"`
	Appfw string `json:"appfw,omitempty"`
	Responder string `json:"responder,omitempty"`
	Agee string `json:"agee,omitempty"`
	Nsxn string `json:"nsxn,omitempty"`
	Modelid string `json:"modelid,omitempty"`
	Push string `json:"push,omitempty"`
	Appflow string `json:"appflow,omitempty"`
	Cloudbridge string `json:"cloudbridge,omitempty"`
	Cloudbridgeappliance string `json:"cloudbridgeappliance,omitempty"`
	Cloudextenderappliance string `json:"cloudextenderappliance,omitempty"`
	Isis string `json:"isis,omitempty"`
	Cluster string `json:"cluster,omitempty"`
	Ch string `json:"ch,omitempty"`
	Appqoe string `json:"appqoe,omitempty"`
	Appflowica string `json:"appflowica,omitempty"`
	Isstandardlic string `json:"isstandardlic,omitempty"`
	Isenterpriselic string `json:"isenterpriselic,omitempty"`
	Isplatinumlic string `json:"isplatinumlic,omitempty"`
	Issgwylic string `json:"issgwylic,omitempty"`
	Isswglic string `json:"isswglic,omitempty"`
	Feo string `json:"feo,omitempty"`
	Lsn string `json:"lsn,omitempty"`
	Licensingmode string `json:"licensingmode,omitempty"`
	Cloudsubscriptionimage string `json:"cloudsubscriptionimage,omitempty"`
	Daystoexpiration string `json:"daystoexpiration,omitempty"`
	Rdpproxy string `json:"rdpproxy,omitempty"`
	Rep string `json:"rep,omitempty"`
	Urlfiltering string `json:"urlfiltering,omitempty"`
	Videooptimization string `json:"videooptimization,omitempty"`
	Forwardproxy string `json:"forwardproxy,omitempty"`
	Sslinterception string `json:"sslinterception,omitempty"`
	Remotecontentinspection string `json:"remotecontentinspection,omitempty"`
	Adaptivetcp string `json:"adaptivetcp,omitempty"`
	Cqa string `json:"cqa,omitempty"`
	Bot string `json:"bot,omitempty"`
	Apigateway string `json:"apigateway,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
