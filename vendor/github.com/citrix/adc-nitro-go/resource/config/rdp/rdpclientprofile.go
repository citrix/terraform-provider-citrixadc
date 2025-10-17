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

package rdp

/**
* Configuration for RDP clientprofile resource.
*/
type Rdpclientprofile struct {
	/**
	* The name of the rdp profile
	*/
	Name string `json:"name,omitempty"`
	/**
	* This setting determines whether the RDP parameters supplied in the vpn url override those specified in the RDP profile.
	*/
	Rdpurloverride string `json:"rdpurloverride,omitempty"`
	/**
	* This setting corresponds to the Clipboard check box on the Local Resources tab under Options in RDC.
	*/
	Redirectclipboard string `json:"redirectclipboard,omitempty"`
	/**
	* This setting corresponds to the selections for Drives under More on the Local Resources tab under Options in RDC.
	*/
	Redirectdrives string `json:"redirectdrives,omitempty"`
	/**
	* This setting corresponds to the selection in the Printers check box on the Local Resources tab under Options in RDC.
	*/
	Redirectprinters string `json:"redirectprinters,omitempty"`
	/**
	* This setting corresponds to the selections for comports under More on the Local Resources tab under Options in RDC.
	*/
	Redirectcomports string `json:"redirectcomports,omitempty"`
	/**
	* This setting corresponds to the selections for pnpdevices under More on the Local Resources tab under Options in RDC.
	*/
	Redirectpnpdevices string `json:"redirectpnpdevices,omitempty"`
	/**
	* This setting corresponds to the selection in the Keyboard drop-down list on the Local Resources tab under Options in RDC.
	*/
	Keyboardhook string `json:"keyboardhook,omitempty"`
	/**
	* This setting corresponds to the selections in the Remote audio area on the Local Resources tab under Options in RDC.
	*/
	Audiocapturemode string `json:"audiocapturemode,omitempty"`
	/**
	* This setting determines if Remote Desktop Connection (RDC) will use RDP efficient multimedia streaming for video playback.
	*/
	Videoplaybackmode string `json:"videoplaybackmode,omitempty"`
	/**
	* Enable/Disable Multiple Monitor Support for Remote Desktop Connection (RDC).
	*/
	Multimonitorsupport string `json:"multimonitorsupport,omitempty"`
	/**
	* RDP cookie validity period. RDP cookie validity time is applicable for new connection and also for any re-connection that might happen, mostly due to network disruption or during fail-over.
	*/
	Rdpcookievalidity *int `json:"rdpcookievalidity,omitempty"`
	/**
	* Add username in rdp file.
	*/
	Addusernameinrdpfile string `json:"addusernameinrdpfile,omitempty"`
	/**
	* RDP file name to be sent to End User
	*/
	Rdpfilename string `json:"rdpfilename,omitempty"`
	/**
	* Fully-qualified domain name (FQDN) of the RDP Listener.
	*/
	Rdphost string `json:"rdphost,omitempty"`
	/**
	* IP address (or) Fully-qualified domain name(FQDN) of the RDP Listener with the port in the format IP:Port (or) FQDN:Port
	*/
	Rdplistener string `json:"rdplistener,omitempty"`
	/**
	* Option for RDP custom parameters settings (if any). Custom params needs to be separated by '&'
	*/
	Rdpcustomparams string `json:"rdpcustomparams,omitempty"`
	/**
	* Pre shared key value
	*/
	Psk string `json:"psk,omitempty"`
	/**
	* Will generate unique filename everytime rdp file is downloaded by appending output of time() function in the format <rdpfileName>_<time>.rdp. This tries to avoid the pop-up for replacement of existing rdp file during each rdp connection launch, hence providing better end-user experience.
	*/
	Randomizerdpfilename string `json:"randomizerdpfilename,omitempty"`
	/**
	* Citrix Gateway allows the configuration of rdpLinkAttribute parameter which can be used to fetch a list of RDP servers(IP/FQDN) that a user can access, from an Authentication server attribute(Example: LDAP, SAML). Based on the list received, the RDP links will be generated and displayed to the user.
		Note: The Attribute mentioned in the rdpLinkAttribute should be fetched through corresponding authentication method.
	*/
	Rdplinkattribute string `json:"rdplinkattribute,omitempty"`
	/**
	* This setting determines whether RDC launch is initiated by the valid client IP
	*/
	Rdpvalidateclientip string `json:"rdpvalidateclientip,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
