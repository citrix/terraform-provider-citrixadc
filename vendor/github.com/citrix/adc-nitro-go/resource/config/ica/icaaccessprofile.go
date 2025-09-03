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

package ica

/**
* Configuration for ica accessprofile resource.
*/
type Icaaccessprofile struct {
	/**
	* Name for the ICA accessprofile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and
		the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the ICA accessprofile is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my ica accessprofile" or 'my ica accessprofile').
		Each of the features can be configured as DEFAULT/DISABLED.
		Here, DISABLED means that the policy settings on the backend XenApp/XenDesktop server are overridden and the Citrix ADC makes the decision to deny access. Whereas DEFAULT means that the Citrix ADC allows the request to reach the XenApp/XenDesktop that takes the decision to allow/deny access based on the policy configured on it. For example, if ClientAudioRedirection is enabled on the backend XenApp/XenDesktop server, and the configured profile has ClientAudioRedirection as DISABLED, the Citrix ADC makes the decision to deny the request irrespective of the configuration on the backend. If the configured profile has ClientAudioRedirection as DEFAULT, then the Citrix ADC forwards the requests to the backend XenApp/XenDesktop server.It then makes the decision to allow/deny access based on the policy configured on it.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Allow Default access/Disable automatic connection of LPT ports from the client when the user logs on
	*/
	Connectclientlptports string `json:"connectclientlptports,omitempty"`
	/**
	* Allow Default access/Disable applications hosted on the server to play sounds through a sound device installed on the client computer, also allows or prevents users to record audio input
	*/
	Clientaudioredirection string `json:"clientaudioredirection,omitempty"`
	/**
	* Allow Default access/Disable file/data sharing via the Receiver for HTML5
	*/
	Localremotedatasharing string `json:"localremotedatasharing,omitempty"`
	/**
	* Allow Default access/Disable the clipboard on the client device to be mapped to the clipboard on the server
	*/
	Clientclipboardredirection string `json:"clientclipboardredirection,omitempty"`
	/**
	* Allow Default access/Disable COM port redirection to and from the client
	*/
	Clientcomportredirection string `json:"clientcomportredirection,omitempty"`
	/**
	* Allow Default access/Disables drive redirection to and from the client
	*/
	Clientdriveredirection string `json:"clientdriveredirection,omitempty"`
	/**
	* Allow Default access/Disable client printers to be mapped to a server when a user logs on to a session
	*/
	Clientprinterredirection string `json:"clientprinterredirection,omitempty"`
	/**
	* Allow Default access/Disable the multistream feature for the specified users
	*/
	Multistream string `json:"multistream,omitempty"`
	/**
	* Allow Default access/Disable the redirection of USB devices to and from the client
	*/
	Clientusbdriveredirection string `json:"clientusbdriveredirection,omitempty"`
	/**
	* Allow default access or disable TWAIN devices, such as digital cameras or scanners, on the client device from published image processing applications
	*/
	Clienttwaindeviceredirection string `json:"clienttwaindeviceredirection,omitempty"`
	/**
	* Allow default access or disable WIA scanner redirection
	*/
	Wiaredirection string `json:"wiaredirection,omitempty"`
	/**
	* Allow default access or disable drag and drop between client and remote applications and desktops
	*/
	Draganddrop string `json:"draganddrop,omitempty"`
	/**
	* Allow default access or disable smart card redirection. Smart card virtual channel is always allowed in CVAD
	*/
	Smartcardredirection string `json:"smartcardredirection,omitempty"`
	/**
	* Allow default access or disable FIDO2 redirection
	*/
	Fido2redirection string `json:"fido2redirection,omitempty"`

	//------- Read only Parameter ---------;

	Refcnt string `json:"refcnt,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Isdefault string `json:"isdefault,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
