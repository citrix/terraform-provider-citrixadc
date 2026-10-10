---
subcategory: "NS"
---

# Resource: nsconfig_update

The nsconfig_update resource is used to apply the update operation for ns config.


## Example usage

```hcl
resource "citrixadc_nsconfig_update" "tf_nsupdate" {
    ipaddress = "10.0.1.164"
    netmask   = "255.255.255.0"
}  
```


## Argument Reference

* `ipaddress` - (Optional) IP address of the Citrix ADC. Commonly referred to as NSIP address. This parameter is mandatory to bring up the appliance.
* `netmask` - (Optional) Netmask corresponding to the IP address. This parameter is mandatory to bring up the appliance.
* `nsvlan` - (Optional) VLAN (NSVLAN) for the subnet on which the IP address resides.
* `ifnum` - (Optional) Interfaces of the appliances that must be bound to the NSVLAN.
* `tagged` - (Optional) Specifies that the interfaces will be added as 802.1q tagged interfaces. Packets sent on these interface on this VLAN will have an additional 4-byte 802.1q tag which identifies the VLAN. To use 802.1q tagging, the switch connected to the appliance's interfaces must also be configured for tagging. Possible values: [ YES, NO ]
* `httpport` - (Optional) The HTTP ports on the Web server. This allows the system to perform connection off-load for any client request that has a destination port matching one of these configured ports. (Set of integers.)
* `maxconn` - (Optional) The maximum number of connections that will be made from the system to the web server(s) attached to it. The value entered here is applied globally to all attached servers. `0` means unlimited.
* `maxreq` - (Optional) The maximum number of requests that the system can pass on a particular connection between the system and a server attached to it. Setting this value to `0` allows an unlimited number of requests to be passed.
* `cip` - (Optional) The option to control (enable or disable) the insertion of the actual client IP address into the HTTP header request passed from the client to one, some, or all servers attached to the system. Possible values: [ ENABLED, DISABLED ]
* `cipheader` - (Optional) The text that will be used as the client IP header.
* `cookieversion` - (Optional) The version of the cookie inserted by the system. Possible values: [ 0, 1 ]
* `securecookie` - (Optional) Enable or disable the secure flag for the persistence cookie. Possible values: [ ENABLED, DISABLED ]
* `pmtumin` - (Optional) The minimum Path MTU.
* `pmtutimeout` - (Optional) The Path MTU timeout value in minutes.
* `ftpportrange` - (Optional) Port range configured for FTP services.
* `crportrange` - (Optional) Port range for cache redirection services.
* `timezone` - (Optional) Name of the timezone.
* `grantquotamaxclient` - (Optional) The percentage of shared quota to be granted at a time for maxClient.
* `exclusivequotamaxclient` - (Optional) The percentage of maxClient to be given to PEs.
* `grantquotaspillover` - (Optional) The percentage of shared quota to be granted at a time for spillover.
* `exclusivequotaspillover` - (Optional) The percentage of spillover threshold to be given to PEs.
* `securemanagementtraffic` - (Optional) Enables secure management traffic handling. Possible values: [ ENABLED, DISABLED ]
* `securemanagementtd` - (Optional) Positive integer that identifies the Management traffic domain. If not specified, defaults to 4094.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsconfig_update. It is a random string prefixed with "tf-nsconfig-update"
