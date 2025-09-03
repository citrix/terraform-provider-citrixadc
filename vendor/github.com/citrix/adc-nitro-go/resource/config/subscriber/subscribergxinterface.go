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

package subscriber

/**
* Configuration for Gx interface Parameters resource.
*/
type Subscribergxinterface struct {
	/**
	* Name of the load balancing, or content switching vserver to which the Gx connections are established. The service type of the virtual server must be DIAMETER/SSL_DIAMETER. Mutually exclusive with the service parameter. Therefore, you cannot set both service and the Virtual Server in the Gx Interface.
	*/
	Vserver string `json:"vserver,omitempty"`
	/**
	* Name of DIAMETER/SSL_DIAMETER service corresponding to PCRF to which the Gx connection is established. The service type of the service must be DIAMETER/SSL_DIAMETER. Mutually exclusive with vserver parameter. Therefore, you cannot set both Service and the Virtual Server in the Gx Interface.
	*/
	Service string `json:"service,omitempty"`
	/**
	* PCRF realm is of type DiameterIdentity and contains the realm of PCRF to which the message is to be routed. This is the realm used in Destination-Realm AVP by Citrix ADC Gx client (as a Diameter node).
	*/
	Pcrfrealm string `json:"pcrfrealm,omitempty"`
	/**
	* Set this setting to yes if Citrix ADC needs to Hold pakcets till subscriber session is fetched from PCRF. Else set to NO. By default set to yes. If this setting is set to NO, then till Citrix ADC fetches subscriber from PCRF, default subscriber profile will be applied to this subscriber if configured. If default subscriber profile is also not configured an undef would be raised to expressions which use Subscriber attributes. 
	*/
	Holdonsubscriberabsence string `json:"holdonsubscriberabsence,omitempty"`
	/**
	* q!Time, in seconds, within which the Gx CCR request must complete. If the request does not complete within this time, the request is retransmitted for requestRetryAttempts time. If still reuqest is not complete then default subscriber profile will be applied to this subscriber if configured. If default subscriber profile is also not configured an undef would be raised to expressions which use Subscriber attributes.
		Zero disables the timeout. !
	*/
	Requesttimeout int `json:"requesttimeout,omitempty"`
	/**
	* If the request does not complete within requestTimeout time, the request is retransmitted for requestRetryAttempts time.
	*/
	Requestretryattempts int `json:"requestretryattempts,omitempty"`
	/**
	* q!Idle Time, in seconds, after which the Gx CCR-U request will be sent after any PCRF activity on a session. Any RAR or CCA message resets the timer.
		Zero value disables the idle timeout. !
	*/
	Idlettl int `json:"idlettl,omitempty"`
	/**
	* q!Revalidation Timeout, in seconds, after which the Gx CCR-U request will be sent after any PCRF activity on a session. Any RAR or CCA message resets the timer.
		Zero value disables the idle timeout. !
	*/
	Revalidationtimeout int `json:"revalidationtimeout,omitempty"`
	/**
	* q!Set this setting to yes if Citrix ADC should send DWR packets to PCRF server. When the session is idle, healthcheck timer expires and DWR packets are initiated in order to check that PCRF server is active. By default set to No. !
	*/
	Healthcheck string `json:"healthcheck,omitempty"`
	/**
	* q!Healthcheck timeout, in seconds, after which the DWR will be sent in order to ensure the state of the PCRF server. Any CCR, CCA, RAR or RRA message resets the timer. !
	*/
	Healthcheckttl int `json:"healthcheckttl,omitempty"`
	/**
	* q!Healthcheck request timeout, in seconds, after which the Citrix ADC considers that no CCA packet received to the initiated CCR. After this time Citrix ADC should send again CCR to PCRF server. !
	*/
	Cerrequesttimeout int `json:"cerrequesttimeout,omitempty"`
	/**
	* q!Negative TTL, in seconds, after which the Gx CCR-I request will be resent for sessions that have not been resolved by PCRF due to server being down or no response or failed response. Instead of polling the PCRF server constantly, negative-TTL makes Citrix ADC stick to un-resolved session. Meanwhile Citrix ADC installs a negative session to avoid going to PCRF.
		For Negative Sessions, Netcaler inherits the attributes from default subscriber profile if default subscriber is configured. A default subscriber could be configured as 'add subscriber profile *'. Or these attributes can be inherited from Radius as well if Radius is configued.
		Zero value disables the Negative Sessions. And Citrix ADC does not install Negative sessions even if subscriber session could not be fetched. !
	*/
	Negativettl int `json:"negativettl,omitempty"`
	/**
	* Set this to YES if Citrix ADC should create negative session for Result-Code DIAMETER_LIMITED_SUCCESS (2002) received in CCA-I. If set to NO, regular session is created.
	*/
	Negativettllimitedsuccess string `json:"negativettllimitedsuccess,omitempty"`
	/**
	* Set this setting to YES if needed to purge Subscriber Database in case of Gx failure. By default set to NO. 
	*/
	Purgesdbongxfailure string `json:"purgesdbongxfailure,omitempty"`
	/**
	*  The AVP code in which PCRF sends service path applicable for subscriber.
	*/
	Servicepathavp []int `json:"servicepathavp,omitempty"`
	/**
	*  The vendorid of the AVP in which PCRF sends service path for subscriber.
	*/
	Servicepathvendorid int `json:"servicepathvendorid,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid int `json:"nodeid,omitempty"`

	//------- Read only Parameter ---------;

	Svrstate string `json:"svrstate,omitempty"`
	Identity string `json:"identity,omitempty"`
	Realm string `json:"realm,omitempty"`
	Status string `json:"status,omitempty"`
	Servicepathinfomode string `json:"servicepathinfomode,omitempty"`
	Gxreportingavp1 string `json:"gxreportingavp1,omitempty"`
	Gxreportingavp1vendorid string `json:"gxreportingavp1vendorid,omitempty"`
	Gxreportingavp1type string `json:"gxreportingavp1type,omitempty"`
	Gxreportingavp2 string `json:"gxreportingavp2,omitempty"`
	Gxreportingavp2vendorid string `json:"gxreportingavp2vendorid,omitempty"`
	Gxreportingavp2type string `json:"gxreportingavp2type,omitempty"`
	Gxreportingavp3 string `json:"gxreportingavp3,omitempty"`
	Gxreportingavp3vendorid string `json:"gxreportingavp3vendorid,omitempty"`
	Gxreportingavp3type string `json:"gxreportingavp3type,omitempty"`
	Gxreportingavp4 string `json:"gxreportingavp4,omitempty"`
	Gxreportingavp4vendorid string `json:"gxreportingavp4vendorid,omitempty"`
	Gxreportingavp4type string `json:"gxreportingavp4type,omitempty"`
	Gxreportingavp5 string `json:"gxreportingavp5,omitempty"`
	Gxreportingavp5vendorid string `json:"gxreportingavp5vendorid,omitempty"`
	Gxreportingavp5type string `json:"gxreportingavp5type,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
