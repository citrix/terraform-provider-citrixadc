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

package dns

/**
* Configuration for dns key resource.
*/
type Dnskey struct {
	/**
	* Name of the public-private key pair to publish in the zone.
	*/
	Keyname string `json:"keyname,omitempty"`
	/**
	* File name of the public key.
	*/
	Publickey string `json:"publickey,omitempty"`
	/**
	* File name of the private key.
	*/
	Privatekey string `json:"privatekey,omitempty"`
	/**
	* Time period for which to consider the key valid, after the key is used to sign a zone.
	*/
	Expires uint32 `json:"expires,omitempty"`
	/**
	* Units for the expiry period.
	*/
	Units1 string `json:"units1,omitempty"`
	/**
	* Time at which to generate notification of key expiration, specified as number of days, hours, or minutes before expiry. Must be less than the expiry period. The notification is an SNMP trap sent to an SNMP manager. To enable the appliance to send the trap, enable the DNSKEY-EXPIRY SNMP alarm.
	*/
	Notificationperiod uint32 `json:"notificationperiod,omitempty"`
	/**
	* Units for the notification period.
	*/
	Units2 string `json:"units2,omitempty"`
	/**
	* Time to Live (TTL), in seconds, for the DNSKEY resource record created in the zone. TTL is the time for which the record must be cached by the DNS proxies. If the TTL is not specified, either the DNS zone's minimum TTL or the default value of 3600 is used.
	*/
	Ttl uint64 `json:"ttl,omitempty"`
	/**
	* Passphrase for reading the encrypted public/private DNS keys
	*/
	Password string `json:"password,omitempty"`
	/**
	* Name of the zone for which to create a key.
	*/
	Zonename string `json:"zonename,omitempty"`
	/**
	* Type of key to create.
	*/
	Keytype string `json:"keytype,omitempty"`
	/**
	* Algorithm to generate for zone signing.
	*/
	Algorithm string `json:"algorithm,omitempty"`
	/**
	* Size of the key, in bits.
	*/
	Keysize uint32 `json:"keysize,omitempty"`
	/**
	* Common prefix for the names of the generated public and private key files and the Delegation Signer (DS) resource record. During key generation, the .key, .private, and .ds suffixes are appended automatically to the file name prefix to produce the names of the public key, the private key, and the DS record, respectively.
	*/
	Filenameprefix string `json:"filenameprefix,omitempty"`
	/**
	* URL (protocol, host, path, and file name) from where the DNS key file will be imported. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access. This is a mandatory argument
	*/
	Src string `json:"src,omitempty"`

}
