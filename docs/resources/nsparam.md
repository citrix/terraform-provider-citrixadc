---
subcategory: "NS"
---

# Resource: nsparam

The nsparam resource is used to set parameters to the target ADC.


## Example usage

```hcl
resource "citrixadc_nsparam" "tf_nsparam" {
  maxconn = 10
  useproxyport = "DISABLED"
}
```


## Argument Reference

* `maxconn` - (Optional) 
* `maxreq` - (Optional) 
* `cip` - (Optional) Enable or disable the insertion of the actual client IP address into the HTTP header request passed from the client to one, some, or all servers attached to the system. The passed address can then be accessed through a minor modification to the server. * If the CIP header is specified, it will be used as the client IP header. * If the CIP header is not specified, the value that has been set will be used as the client IP header. Possible values: [ ENABLED, DISABLED ]
* `cipheader` - (Optional) Text that will be used as the client IP address header.
* `cookieversion` - (Optional) Version of the cookie inserted by the system. Possible values: [ 0, 1 ]
* `securecookie` - (Optional) Enable or disable secure flag for persistence cookie. Possible values: [ ENABLED, DISABLED ]
* `pmtumin` - (Optional) 
* `pmtutimeout` - (Optional) Interval, in minutes, for flushing the PMTU entries.
* `ftpportrange` - (Optional) 
* `crportrange` - (Optional) Port range for cache redirection services.
* `timezone` - (Optional) Time zone for the Citrix ADC. Name of the time zone should be specified as argument.
* `grantquotamaxclient` - (Optional) Percentage of shared quota to be granted at a time for maxClient.
* `exclusivequotamaxclient` - (Optional) Percentage of maxClient to be given to PEs.
* `grantquotaspillover` - (Optional) Percentage of shared quota to be granted at a time for spillover.
* `exclusivequotaspillover` - (Optional) Percentage of maximum limit to be given to PEs.
* `useproxyport` - (Optional) Enable/Disable use_proxy_port setting. Possible values: [ ENABLED, DISABLED ]
* `internaluserlogin` - (Optional) Enables/disables the internal user from logging in to the appliance. Before disabling internal user login, you must have key-based authentication set up on the appliance. The file name for the key pair must be "ns_comm_key". Possible values: [ ENABLED, DISABLED ]
* `aftpallowrandomsourceport` - (Optional) Allow the FTP server to come from a random source port for active FTP data connections. Possible values: [ ENABLED, DISABLED ]
* `tcpcip` - (Optional) Enable or disable the insertion of the client TCP/IP header in TCP payload passed from the client to one, some, or all servers attached to the system. The passed address can then be accessed through a minor modification to the server. Possible values: [ ENABLED, DISABLED ]
* `servicepathingressvlan` - (Optional) VLAN on which the subscriber traffic arrives on the appliance.
* `mgmthttpport` - (Optional) This allow the configuration of management HTTP port.
* `mgmthttpsport` - (Optional) This allows the configuration of management HTTPS port.
* `proxyprotocol` - (Optional) Disable/Enable v1 or v2 proxy protocol header for client info insertion. Possible values: [ ENABLED, DISABLED ]
* `advancedanalyticsstats` - (Optional) Disable/Enable advanace analytics stats. Possible values: [ ENABLED, DISABLED ]
* `icaports` - (Optional) The ICA ports on the Web server. This allows the system to perform connection off-load for any client request that has a destination port matching one of these configured ports.
* `secureicaports` - (Optional) The Secure ICA ports on the Web server. This allows the system to perform connection off-load for any client request that has a destination port matching one of these configured ports.
* `ipttl` - (Optional) Set the IP Time to Live (TTL) and Hop Limit value for all outgoing packets from Citrix ADC.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsparam. It is a random string prefixed with "tf-nsparam-"
