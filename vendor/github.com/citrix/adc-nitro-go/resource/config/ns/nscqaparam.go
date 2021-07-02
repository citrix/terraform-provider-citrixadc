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
* Configuration for cqaparam resource.
*/
type Nscqaparam struct {
	/**
	* HARQ retransmission delay (in ms).
	*/
	Harqretxdelay uint32 `json:"harqretxdelay,omitempty"`
	/**
	* Name of the network label.
	*/
	Net1label string `json:"net1label,omitempty"`
	/**
	* MIN RTT (in ms) for the first network.
	*/
	Minrttnet1 uint32 `json:"minrttnet1,omitempty"`
	/**
	* coefficients values for Label1.
	*/
	Lr1coeflist string `json:"lr1coeflist,omitempty"`
	/**
	* Probability threshold values for LR model to differentiate between NET1 and reset(NET2 and NET3).
	*/
	Lr1probthresh float64 `json:"lr1probthresh,omitempty"`
	/**
	* Three congestion level scores limits corresponding to None, Low, Medium.
	*/
	Net1cclscale string `json:"net1cclscale,omitempty"`
	/**
	* Three signal quality level scores limits corresponding to Excellent, Good, Fair.
	*/
	Net1csqscale string `json:"net1csqscale,omitempty"`
	/**
	* Connection quality ranking Log coefficients of network 1.
	*/
	Net1logcoef string `json:"net1logcoef,omitempty"`
	/**
	* Name of the network label 2.
	*/
	Net2label string `json:"net2label,omitempty"`
	/**
	* MIN RTT (in ms) for the second network.
	*/
	Minrttnet2 uint32 `json:"minrttnet2,omitempty"`
	/**
	* coefficients values for Label 2.
	*/
	Lr2coeflist string `json:"lr2coeflist,omitempty"`
	/**
	* Probability threshold values for LR model to differentiate between NET2 and NET3.
	*/
	Lr2probthresh float64 `json:"lr2probthresh,omitempty"`
	/**
	* Three congestion level scores limits corresponding to None, Low, Medium.
	*/
	Net2cclscale string `json:"net2cclscale,omitempty"`
	/**
	* Three signal quality level scores limits corresponding to Excellent, Good, Fair.
	*/
	Net2csqscale string `json:"net2csqscale,omitempty"`
	/**
	* Connnection quality ranking Log coefficients of network 2.
	*/
	Net2logcoef string `json:"net2logcoef,omitempty"`
	/**
	* Name of the network label 3.
	*/
	Net3label string `json:"net3label,omitempty"`
	/**
	* MIN RTT (in ms) for the third network.
	*/
	Minrttnet3 uint32 `json:"minrttnet3,omitempty"`
	/**
	* Three congestion level scores limits corresponding to None, Low, Medium.
	*/
	Net3cclscale string `json:"net3cclscale,omitempty"`
	/**
	* Three signal quality level scores limits corresponding to Excellent, Good, Fair.
	*/
	Net3csqscale string `json:"net3csqscale,omitempty"`
	/**
	* Connection quality ranking Log coefficients of network 3.
	*/
	Net3logcoef string `json:"net3logcoef,omitempty"`

}
