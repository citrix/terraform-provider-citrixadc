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
* Configuration for expr evaluation resource.
*/
type Policyevaluation struct {
	/**
	* Expression string. For example: http.req.body(100).contains("this").
	*/
	Expression string `json:"expression,omitempty"`
	/**
	* Rewrite action name. Supported rewrite action types are:
		-delete
		-delete_all
		-delete_http_header
		-insert_after
		-insert_after_all
		-insert_before
		-insert_before_all
		-insert_http_header
		-replace
		-replace_all
	*/
	Action string `json:"action,omitempty"`
	/**
	* Indicates request or response input packet
	*/
	Type string `json:"type,omitempty"`
	/**
	* Text representation of input packet.
	*/
	Input string `json:"input,omitempty"`

	//------- Read only Parameter ---------;

	Pitmodifiedinputdata string `json:"pitmodifiedinputdata,omitempty"`
	Pitboolresult string `json:"pitboolresult,omitempty"`
	Pitnumresult string `json:"pitnumresult,omitempty"`
	Pitdoubleresult string `json:"pitdoubleresult,omitempty"`
	Pitulongresult string `json:"pitulongresult,omitempty"`
	Pitrefresult string `json:"pitrefresult,omitempty"`
	Pitoffsetresult string `json:"pitoffsetresult,omitempty"`
	Pitoffsetresultlen string `json:"pitoffsetresultlen,omitempty"`
	Istruncatedrefresult string `json:"istruncatedrefresult,omitempty"`
	Pitboolevaltime string `json:"pitboolevaltime,omitempty"`
	Pitnumevaltime string `json:"pitnumevaltime,omitempty"`
	Pitdoubleevaltime string `json:"pitdoubleevaltime,omitempty"`
	Pitulongevaltime string `json:"pitulongevaltime,omitempty"`
	Pitrefevaltime string `json:"pitrefevaltime,omitempty"`
	Pitoffsetevaltime string `json:"pitoffsetevaltime,omitempty"`
	Pitactionevaltime string `json:"pitactionevaltime,omitempty"`
	Pitoperationperformerarray string `json:"pitoperationperformerarray,omitempty"`
	Pitoldoffsetarray string `json:"pitoldoffsetarray,omitempty"`
	Pitnewoffsetarray string `json:"pitnewoffsetarray,omitempty"`
	Pitoffsetlengtharray string `json:"pitoffsetlengtharray,omitempty"`
	Pitoffsetnewlengtharray string `json:"pitoffsetnewlengtharray,omitempty"`
	Pitboolerrorresult string `json:"pitboolerrorresult,omitempty"`
	Pitnumerrorresult string `json:"pitnumerrorresult,omitempty"`
	Pitdoubleerrorresult string `json:"pitdoubleerrorresult,omitempty"`
	Pitulongerrorresult string `json:"pitulongerrorresult,omitempty"`
	Pitreferrorresult string `json:"pitreferrorresult,omitempty"`
	Pitoffseterrorresult string `json:"pitoffseterrorresult,omitempty"`
	Pitactionerrorresult string `json:"pitactionerrorresult,omitempty"`

}
